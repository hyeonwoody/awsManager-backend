package user_useCase

import (
	user "awsManager/api/user/cmd"
	dto "awsManager/api/user/cmd/useCase/dto"
)

type IUserProjectFacade interface {
	FindNextIndex(projectName string) (uint, error)
	CreateUser(input dto.CreateUserCommand) (*user.Model, error)
}
