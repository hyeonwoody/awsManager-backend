package ec2_domain

import (
	dto "awsManager/api/ec2/cmd/dto"
	ec2 "awsManager/api/ec2/cmd/model"
)

type IService interface {
	Create(command *dto.CreateCommand) (*ec2.Model, error)
}
