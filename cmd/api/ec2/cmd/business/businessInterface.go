package ec2_business

import (
	dto "awsManager/api/ec2/cmd/business/dto"
	domainDto "awsManager/api/ec2/cmd/domain/dto"
)

type IBusiness interface {
	Create(command *domainDto.CreateCommand) (*dto.Ec2Instance, error)
	InitWithPublicIp(command *domainDto.InitWithPublicIpCommand) error
	Delete(command *domainDto.DeleteCommand) error
}
