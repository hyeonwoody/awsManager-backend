package user_useCase

import (
	dto "awsManager/api/user/cmd/application/useCase/dto/in"
	user "awsManager/api/user/cmd/model"
)

type IUserProjectFacade interface {
	FindNextIndex(projectName string) (uint, error)
	CreateUser(input dto.CreateUserCommand) (*user.Model, error)
}
