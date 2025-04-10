package ec2_domain

import (
	dto "awsManager/api/ec2/cmd/domain/dto"
	ec2 "awsManager/api/ec2/cmd/model"
	ec2Model "awsManager/api/ec2/cmd/model"
)

type IService interface {
	Create(command *dto.CreateCommand) (*ec2.Model, error)
	DeleteExist(command *dto.DeleteCommand) error
	FindByInstanceId(instanceId *string) (*ec2.Model, error)
	AttachEbsVolume(command *dto.AttachEbsVolumeCommand) error
	AddMemory(command *dto.AddMemoryCommand) (*ec2.Model, error)
	InstallDocker(command *dto.InstallCommand) (*ec2.Model, error)
	InstallDockerNginx(command *dto.InstallDockerNginxCommand) (*ec2Model.Model, error)
	InstallGoAgent(command *dto.InstallGoAgentCommand) (*ec2Model.Model, error)
	InstallDockerGoAgent(command *dto.InstallDockerGoAgentCommand) (*ec2Model.Model, error)
	InstallGoServer(command *dto.InstallGocdCommand) (*ec2Model.Model, error)
	GetProxyNginxIp() string
	GetMyIp() string
	AddInboundRule(accessKey, secretAccessKey, keyName, publicIp *string) (*string, error)
}
