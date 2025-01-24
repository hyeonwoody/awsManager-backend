package ec2_useCase

import (
	useCaseDto "awsManager/api/ec2/cmd/application/useCase/dto/in"
	ec2Domain "awsManager/api/ec2/cmd/domain"
	ec2DomainDto "awsManager/api/ec2/cmd/domain/dto"
	projectDomain "awsManager/api/project/cmd/domain"
	userDomain "awsManager/api/user/cmd/domain"
)

type Ec2UserProjectFacade struct {
	ec2Svc     ec2Domain.IService
	userSvc    userDomain.IService
	projectSvc projectDomain.IService
}

func NewEc2UserProjectFacade(ec2Svc ec2Domain.IService, userSvc userDomain.IService, projectSvc projectDomain.IService) *Ec2UserProjectFacade {
	return &Ec2UserProjectFacade{
		ec2Svc:     ec2Svc,
		userSvc:    userSvc,
		projectSvc: projectSvc,
	}
}

func (f *Ec2UserProjectFacade) Create(command *useCaseDto.CreateEc2Command) (interface{}, error) {
	project, err := f.projectSvc.FindByName(command.ProjectName)
	if err != nil {
		return nil, err
	}
	user, err := f.userSvc.FindByProjectIdAndKey(project.Id, command.KeyNumber)
	if err != nil {
		return nil, err
	}
	// if user.Ec2InstanceId != "" {
	// 	return nil, fmt.Errorf("user already have instance")
	// }
	//ec2Err :=
	f.ec2Svc.DeleteExist(ec2DomainDto.DeleteCommandFrom(project.Name, user.AccessKey, user.SecretAccessKey, project.Id, user.KeyNumber))
	// if ec2Err != nil {
	// 	return nil, err
	// }
	ec2, err := f.ec2Svc.Create(ec2DomainDto.CreateCommandFrom(project.Name, command.Ami, command.InstanceType, user.AccessKey, user.SecretAccessKey, project.Id, user.KeyNumber))
	if err != nil {
		return nil, err
	}
	user.Ec2InstanceId = ec2.InstanceId
	f.userSvc.Save(user)

	f.ec2Svc.Init(ec2DomainDto.InitWithPublicIpCommandFrom(ec2.PublicIp, project.Name, user.KeyNumber))

	return ec2, nil
}

func (f *Ec2UserProjectFacade) AttachEbsVolume(input *useCaseDto.AttachEbsVolumeCommand) (interface{}, error) {
	ec2, err := f.ec2Svc.FindByInstanceId(&input.InstanceId)
	project, err := f.projectSvc.Read(ec2.ProjectId)
	user, err := f.userSvc.FindByProjectIdAndKey(ec2.ProjectId, ec2.KeyNumber)
	if err != nil {
		return nil, err
	}

	f.ec2Svc.AttachEbsVolume(ec2DomainDto.AttachEbsVolumeCommandFrom(user.AccessKey, user.SecretAccessKey, project.Name, ec2))
	return nil, nil
}

func (f *Ec2UserProjectFacade) AddMemory(input *useCaseDto.InitEc2Command) (interface{}, error) {
	ec2, _ := f.ec2Svc.FindByInstanceId(&input.InstanceId)
	project, err := f.projectSvc.Read(ec2.ProjectId)

	f.ec2Svc.AddMemory(ec2DomainDto.AddMemoryCommandFrom(ec2.PublicIp, project.Name, ec2.KeyNumber))
	if err != nil {
		return nil, err
	}
	return nil, nil
}
