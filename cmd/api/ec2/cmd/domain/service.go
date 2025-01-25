package ec2_domain

import (
	ec2Businses "awsManager/api/ec2/cmd/business"
	businessDto "awsManager/api/ec2/cmd/business/dto"
	dto "awsManager/api/ec2/cmd/domain/dto"
	ec2Infrastructure "awsManager/api/ec2/cmd/infrastructure"
	ec2Model "awsManager/api/ec2/cmd/model"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Service struct {
	sdkBiz ec2Businses.SdkBusiness
	cliBiz ec2Businses.CliBusiness
	repo   ec2Infrastructure.IRepository
}

func NewService(sdkBiz ec2Businses.SdkBusiness, cliBiz ec2Businses.CliBusiness, repo ec2Infrastructure.IRepository) *Service {
	return &Service{sdkBiz: sdkBiz, cliBiz: cliBiz, repo: repo}
}

func (s *Service) DeleteExist(command *dto.DeleteCommand) error {
	//ctx := s.sdkBiz.GetContext()
	s.sdkBiz.Delete(command)
	err := s.repo.DeleteByIdAndKeyNumber(command.ProjectId, command.KeyNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Create(command *dto.CreateCommand) (*ec2Model.Model, error) {
	client, err := s.sdkBiz.GetAsyncClient(&command.AccessKey, &command.SecretAccessKey)
	if err != nil {
		return nil, err
	}
	ec2Instance, err := s.sdkBiz.Create(command, client)
	if err != nil {
		return nil, err
	}
	ec2, err := s.repo.Save(businessDto.ModelFrom(command, ec2Instance))
	if err != nil {
		return nil, err
	}
	return ec2, nil
}

func (s *Service) AddMemory(command *dto.AddMemoryCommand) (*ec2Model.Model, error) {
	s.cliBiz.AddMemory(command)
	return nil, nil
}

func (s *Service) FindByInstanceId(instanceId *string) (*ec2Model.Model, error) {
	ec2, err := s.repo.FindByInstanceId(instanceId)
	if err != nil {
		return nil, err
	}
	return ec2, nil
}

func (s *Service) AttachEbsVolume(command *dto.AttachEbsVolumeCommand) error {
	client, _ := s.sdkBiz.GetAsyncClient(&command.AccessKey, &command.SecretAccessKey)

	var privateKeyName = command.ProjectName + strconv.Itoa(int(command.Ec2.KeyNumber))
	// var path = string("/dev/zzz")
	// s.cliBiz.MakeDir(dto.CliCommandFrom(
	// 	&privateKeyName,
	// 	&command.Ec2.PublicIp,
	// 	&path))
	availabilityZone, _ := s.sdkBiz.GetEc2AvailibityZone(client, command.Ec2.InstanceId)

	{
		var memoryDeviceName = "xvdm"
		s.attachEbs(client, &availabilityZone, command.Ec2, &memoryDeviceName, 9)
		err := s.cliBiz.MountEbsVolume(dto.CliCommandFrom(
			&privateKeyName,
			&command.Ec2.PublicIp,
			&memoryDeviceName,
		))
		if err != nil {
			//s.dettachEbs()
		}
	}
	{
		var storageDeviceName = "xvdf"
		s.attachEbs(client, &availabilityZone, command.Ec2, &storageDeviceName, 7)
		s.cliBiz.MountEbsVolume(dto.CliCommandFrom(
			&privateKeyName,
			&command.Ec2.PublicIp,
			&storageDeviceName,
		))
	}

	_, saveErr := s.repo.Save(businessDto.ModelFromAttachVolume(command.Ec2))
	if saveErr != nil {
		return saveErr
	}
	return nil
}

func (s *Service) attachEbs(client *ec2.Client, availabilityZone *string, ec2 *ec2Model.Model, deviceName *string, size uint) error {
	volumeId, err := s.sdkBiz.CreateEbsVolume(client, availabilityZone, size)
	if err != nil {
		return err
	}
	attachErr := s.sdkBiz.AttachEbsVolume(client, &ec2.InstanceId, volumeId, deviceName)
	if attachErr != nil {
		s.sdkBiz.DetachEbsVolume(client, volumeId)
		return attachErr
	}
	return nil
}

func (s *Service) InstallDocker(command *dto.InstallCommand) (*ec2Model.Model, error) {
	s.cliBiz.InstallDocker(command)
	return nil, nil
}

func (s *Service) InstallDockerNginx(command *dto.InstallCommand) (*ec2Model.Model, error) {
	s.cliBiz.InstallDockerNginx(command)
	return nil, nil
}

func (s *Service) InstallDockerGoAgent(command *dto.InstallCommand) (*ec2Model.Model, error) {
	s.cliBiz.InstallDockerGoAgent(command)
	return nil, nil
}
