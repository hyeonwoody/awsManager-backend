package ec2_infrastructure

import (
	ec2 "awsManager/api/ec2/cmd/model"
)

type IRepository interface {
	Save(ec2 *ec2.Model) (*ec2.Model, error)
	DeleteByIdAndKeyNumber(projectId uint, KeyNumber uint) error
	FindByInstanceId(instanceId *string) (*ec2.Model, error)
	FindByProjectIdAndKey(projectId, keyNumber uint) (*ec2.Model, error)
	// FindAll() ([]User, error)
	// Update(user *User) error
	// DeleteById(id uint) (*User, error)
}
