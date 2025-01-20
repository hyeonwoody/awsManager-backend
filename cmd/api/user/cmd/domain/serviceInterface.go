package domain

import (
	user "awsManager/api/user/cmd/model"
)

type IService interface {
	FindNextIndex(projectId uint) uint
	Create(projectId, keyNubmber uint, password, accessKey, secretAccessKey string) (*user.Model, error)
}
