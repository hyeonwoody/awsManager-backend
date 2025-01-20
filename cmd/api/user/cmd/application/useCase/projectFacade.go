package user_useCase

import (
	project_domain "awsManager/api/project/cmd/domain"
	dto "awsManager/api/user/cmd/application/useCase/dto/in"
	user_domain "awsManager/api/user/cmd/domain"
	user "awsManager/api/user/cmd/model"
	"fmt"
)

type UserProjectFacade struct {
	userService    user_domain.IService
	projectService project_domain.IService
}

func NewUserProjectFacade(userSvc user_domain.IService, projectSvc project_domain.IService) *UserProjectFacade {
	return &UserProjectFacade{
		userService:    userSvc,
		projectService: projectSvc,
	}
}

func (f *UserProjectFacade) FindNextIndex(projectName string) (uint, error) {
	project, err := f.projectService.FindByName(projectName)
	if err != nil {
		return 0, err
	}
	nextIndex := f.userService.FindNextIndex(project.Id)
	return nextIndex, nil
}

func (f *UserProjectFacade) CreateUser(input dto.CreateUserCommand) (*user.Model, error) {
	project, err := f.projectService.FindByName(input.ProjectName)
	if err != nil {
		return nil, fmt.Errorf("failed to find Project: %w", err)
	}
	createdUser, err := f.userService.Create(project.Id, input.KeyNumber, input.Password, input.AccessKey, input.AccessKey)
	return createdUser, err
}
