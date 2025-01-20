package user_useCase

import (
	project "awsManager/api/project/cmd"
	user "awsManager/api/user/cmd"
	dto "awsManager/api/user/cmd/useCase/dto"
)

type UserProjectFacade struct {
	userService    user.IService
	projectService project.IService
}

func NewUserProjectFacade(userSvc user.IService, projectSvc project.IService) *UserProjectFacade {
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
		return nil, err
	}
	createdUser, err := f.userService.Create(project.Id, input.KeyNumber, input.Password, input.AccessKey, input.AccessKey)
	return createdUser, err
}
