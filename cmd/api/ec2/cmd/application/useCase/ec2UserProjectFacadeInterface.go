package ec2_useCase

import (
	useCaseDto "awsManager/api/ec2/cmd/application/useCase/dto/in"
)

type IEc2UserProjectFacade interface {
	Create(input *useCaseDto.CreateEc2Command) (interface{}, error)
	AddMemory(input *useCaseDto.InitEc2Command) (interface{}, error)
	AttachEbsVolume(input *useCaseDto.AttachEbsVolumeCommand) (interface{}, error)
	InstallDocker(input *useCaseDto.InstallCommand) (interface{}, error)
	InstallDockerNginx(input *useCaseDto.InstallCommand) (interface{}, error)
	InstallGoAgent(input *useCaseDto.InstallCommand) (interface{}, error)
	InstallDockerGoAgent(input *useCaseDto.InstallCommand) (interface{}, error)
	InstallGoServer(input *useCaseDto.InstallCommand) (interface{}, error)
}
