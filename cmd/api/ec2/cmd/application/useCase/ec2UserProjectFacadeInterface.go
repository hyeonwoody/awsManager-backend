package ec2_useCase

import (
	useCaseDto "awsManager/api/ec2/cmd/application/useCase/dto/in"
)

type IEc2UserProjectFacade interface {
	Create(input *useCaseDto.CreateEc2Command) (interface{}, error)
	Init(input *useCaseDto.InitEc2Command) (interface{}, error)
}
