package dependecyinjection

import (
	project "awsManager/api/project/cmd"
	subProject_domain "awsManager/api/project/cmd/subProject/domain"
	subProject_infrastructure "awsManager/api/project/cmd/subProject/infrastructure"
	user "awsManager/api/user/cmd"
	userHandler "awsManager/api/user/cmd/handler"
	userUseCase "awsManager/api/user/cmd/useCase"

	"gorm.io/gorm"
)

type Container struct {
	UserRepository user.IRepository
	UserService    user.IService
	UserHandler    userHandler.IHandler

	ProjectRepository project.IRepository
	ProjectService    project.IService
	ProjectHandler    project.IHandler

	SubProjectRepository subProject_infrastructure.IRepository
	SubProjectService    subProject_domain.IService

	UserProjectFacade userUseCase.IUserProjectFacade
}

func (c *Container) Init(db *gorm.DB) {
	c.UserRepository = user.NewRepository(db)
	c.ProjectRepository = project.NewRepository(db)
	c.SubProjectRepository = subProject_infrastructure.NewRepository(db)

	c.UserService = user.NewService(c.UserRepository)
	c.ProjectService = project.NewService(c.ProjectRepository)
	c.SubProjectService = subProject_domain.NewService(c.SubProjectRepository)

	c.ProjectHandler = project.NewHandler(c.ProjectService, c.SubProjectService)

	c.UserProjectFacade = userUseCase.NewUserProjectFacade(c.UserService, c.ProjectService)
	c.UserHandler = userHandler.NewHandler(c.UserProjectFacade, c.UserService)
}
