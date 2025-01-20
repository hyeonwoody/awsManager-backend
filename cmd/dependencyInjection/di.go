package dependecyinjection

import (
	project_domain "awsManager/api/project/cmd/domain"
	project_infrastructure "awsManager/api/project/cmd/infrastructure"
	project_presentation "awsManager/api/project/cmd/presentation"
	subProject_domain "awsManager/api/project/cmd/subProject/domain"
	subProject_infrastructure "awsManager/api/project/cmd/subProject/infrastructure"
	userUseCase "awsManager/api/user/cmd/application/useCase"

	user_business "awsManager/api/user/cmd/business"
	user_domain "awsManager/api/user/cmd/domain"
	user_infrastructure "awsManager/api/user/cmd/infrastructure"
	user_presentation "awsManager/api/user/cmd/presentation"

	ec2_domain "awsManager/api/ec2/cmd/domain"
	ec2_infrastructure "awsManager/api/ec2/cmd/infrastructure"
	ec2_presentation "awsManager/api/ec2/cmd/presentation"

	"gorm.io/gorm"
)

type Container struct {
	UserRepository user_infrastructure.IRepository
	UserBusiness   user_business.IBusiness
	UserService    user_domain.IService
	UserHandler    user_presentation.IHandler

	ProjectRepository project_infrastructure.IRepository
	ProjectService    project_domain.IService
	ProjectHandler    project_presentation.IHandler

	SubProjectRepository subProject_infrastructure.IRepository
	SubProjectService    subProject_domain.IService

	UserProjectFacade userUseCase.IUserProjectFacade

	Ec2Repository ec2_infrastructure.IRepository
	Ec2Service    ec2_domain.IService
	Ec2Handler    ec2_presentation.IHandler
}

func (c *Container) Init(db *gorm.DB) {
	c.UserRepository = user_infrastructure.NewRepository(db)
	c.ProjectRepository = project_infrastructure.NewRepository(db)
	c.SubProjectRepository = subProject_infrastructure.NewRepository(db)
	c.Ec2Repository = ec2_infrastructure.NewRepository(db)

	c.UserBusiness = user_business.NewBusiness()
	c.UserService = user_domain.NewService(c.UserBusiness, c.UserRepository)
	c.ProjectService = project_domain.NewService(c.ProjectRepository)
	c.SubProjectService = subProject_domain.NewService(c.SubProjectRepository)

	c.ProjectHandler = project_presentation.NewHandler(c.ProjectService, c.SubProjectService)

	c.UserProjectFacade = userUseCase.NewUserProjectFacade(c.UserService, c.ProjectService)
	c.UserHandler = user_presentation.NewHandler(c.UserProjectFacade, c.UserService)

	c.Ec2Service = ec2_domain.NewService(c.Ec2Repository)
	c.Ec2Handler = ec2_presentation.NewHandler(c.Ec2Service)
}
