package domain

import (
	user "awsManager/api/user/cmd/model"
)

type IService interface {
	FindNextIndex(projectId uint) uint
	Create(projectId, keyNubmber uint, projectName, password, accessKey, secretAccessKey string) (*user.Model, error)
	FindByProjectIdAndKey(id uint, keyNumber uint) (*user.Model, error)
	FindInstanceOff(projectId uint) ([]user.Model, error)
	Save(user *user.Model) error
}
