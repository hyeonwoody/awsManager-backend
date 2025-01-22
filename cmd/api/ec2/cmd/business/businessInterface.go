package ec2_business

import (
	dto "awsManager/api/ec2/cmd/dto"
)

type IBusiness interface {
	Create(command *dto.CreateCommand) (*dto.Ec2Instance, error)
	Init(ec2Instance *dto.Ec2Instance) (*dto.Ec2Instance, error)
}
