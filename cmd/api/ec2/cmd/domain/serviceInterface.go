package ec2_domain

import (
	dto "awsManager/api/ec2/cmd/domain/dto"
	ec2 "awsManager/api/ec2/cmd/model"
)

type IService interface {
	Create(command *dto.CreateCommand) (*ec2.Model, error)
	DeleteExist(command *dto.DeleteCommand) error
	AddMemory(command *dto.AddMemoryCommand) (*ec2.Model, error)
	FindByInstanceId(instanceId *string) (*ec2.Model, error)
	AttachEbsVolume(command *dto.AttachEbsVolumeCommand) error
}
