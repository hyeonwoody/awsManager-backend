package ec2_useCase

import (
	dto "awsManager/api/ec2/cmd/application/useCase/dto/in"
	ec2 "awsManager/api/ec2/cmd/model"
)

type IEc2UserProjectFacade interface {
	Init(input *dto.InitEc2Command) (*ec2.Model, error)
}
