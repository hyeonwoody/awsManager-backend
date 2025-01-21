package user_infrastructure

import (
	user "awsManager/api/user/cmd/model"
)

type IRepository interface {
	FindNextIndex(projectId uint) uint
	Save(user *user.Model) error
	FindByProjectIdAndKey(projectId uint, keyNumber uint) (*user.Model, error)
	FindInstanceOff(projectId uint) ([]user.Model, error)
	// Save(user *User) error
	// FindById(id uint) (*User, error)
	// FindAll() ([]User, error)
	// Update(user *User) error
	// DeleteById(id uint) (*User, error)
}
