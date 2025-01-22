package ec2_business

import (
	dto "awsManager/api/ec2/cmd/dto"
)

type IBusiness interface {
	Create(command *dto.CreateCommand) (*dto.Ec2Instance, error)
	InitWithPublicIp(command *dto.InitWithPublicIpCommand) error
	Delete(command *dto.DeleteCommand) error
}
