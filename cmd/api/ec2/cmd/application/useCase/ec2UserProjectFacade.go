package ec2_useCase

import (
	useCaseDto "awsManager/api/ec2/cmd/application/useCase/dto/in"
	ec2Domain "awsManager/api/ec2/cmd/domain"
	ec2DomainDto "awsManager/api/ec2/cmd/domain/dto"
	projectDomain "awsManager/api/project/cmd/domain"
	userDomain "awsManager/api/user/cmd/domain"
	"strconv"
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

func (f *Ec2UserProjectFacade) Create(input *useCaseDto.CreateEc2Command) (interface{}, error) {
	project, err := f.projectSvc.FindByName(input.ProjectName)
	if err != nil {
		return nil, err
	}
	user, err := f.userSvc.FindByProjectIdAndKey(project.Id, input.KeyNumber)
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
	ec2, err := f.ec2Svc.Create(ec2DomainDto.CreateCommandFrom(project.Name, input.Ami, input.InstanceType, user.AccessKey, user.SecretAccessKey, project.Id, user.KeyNumber))
	if err != nil {
		return nil, err
	}

	f.ec2Svc.InstallDocker(ec2DomainDto.InstallCommandFrom(ec2.PublicIp, project.Name, ec2.KeyNumber))

	user.Ec2InstanceId = ec2.InstanceId
	f.userSvc.Save(user)

	//f.ec2Svc.AddMemory(ec2DomainDto.AddMemoryCommandFrom(ec2.PublicIp, project.Name, user.KeyNumber))

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

func (f *Ec2UserProjectFacade) InstallDocker(input *useCaseDto.InstallCommand) (interface{}, error) {
	ec2, _ := f.ec2Svc.FindByInstanceId(&input.InstanceId)
	project, err := f.projectSvc.Read(ec2.ProjectId)

	f.ec2Svc.InstallDocker(ec2DomainDto.InstallCommandFrom(ec2.PublicIp, project.Name, ec2.KeyNumber))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (f *Ec2UserProjectFacade) InstallDockerNginx(input *useCaseDto.InstallCommand) (interface{}, error) {
	ec2, _ := f.ec2Svc.FindByInstanceId(&input.InstanceId)
	project, _ := f.projectSvc.Read(ec2.ProjectId)
	user, err := f.userSvc.FindByProjectIdAndKey(ec2.ProjectId, ec2.KeyNumber)
	f.ec2Svc.InstallDockerNginx(ec2DomainDto.InstallDockerNginxCommandFrom(user.AccessKey, user.SecretAccessKey, ec2.PublicIp, project.Name, ec2.KeyNumber))
	if err != nil {
		return nil, err
	}
	var pcIp = f.ec2Svc.GetMyIp()
	keyName := project.Name + strconv.Itoa(int(ec2.KeyNumber))
	f.ec2Svc.AddInboundRule(&user.AccessKey, &user.SecretAccessKey, &keyName, &pcIp)
	return nil, nil
}

func (f *Ec2UserProjectFacade) InstallGoAgent(input *useCaseDto.InstallCommand) (interface{}, error) {
	ec2, _ := f.ec2Svc.FindByInstanceId(&input.InstanceId)
	project, _ := f.projectSvc.Read(ec2.ProjectId)
	user, err := f.userSvc.FindByProjectIdAndKey(ec2.ProjectId, ec2.KeyNumber)
	goServerIp := f.ec2Svc.GetProxyNginxIp()
	if err != nil {
		return nil, err
	}
	f.ec2Svc.InstallGoAgent(ec2DomainDto.InstallGoAgentCommandFrom(user.AccessKey, user.SecretAccessKey, ec2.PublicIp, project.Name, goServerIp, ec2.KeyNumber))

	shouldReturn, result, err := f.addInboundRuleInBoGocd(&ec2.PublicIp)
	if shouldReturn {
		return result, err
	}
	return nil, nil
}

func (f *Ec2UserProjectFacade) addInboundRuleInBoGocd(publicIp *string) (bool, interface{}, error) {
	gocdUser, err := f.userSvc.FindGocd()
	keyName := "bohemiangocd0"
	_, _ = f.ec2Svc.AddInboundRule(&gocdUser.AccessKey, &gocdUser.SecretAccessKey, publicIp, &keyName)
	if err != nil {
		return true, nil, err
	}
	return false, nil, nil
}

func (f *Ec2UserProjectFacade) InstallDockerGoAgent(input *useCaseDto.InstallCommand) (interface{}, error) {
	ec2, _ := f.ec2Svc.FindByInstanceId(&input.InstanceId)
	project, _ := f.projectSvc.Read(ec2.ProjectId)
	user, err := f.userSvc.FindByProjectIdAndKey(ec2.ProjectId, ec2.KeyNumber)
	goServerIp := f.ec2Svc.GetProxyNginxIp()
	if err != nil {
		return nil, err
	}
	f.ec2Svc.InstallDockerGoAgent(ec2DomainDto.InstallDockerGoAgentCommandFrom(user.AccessKey, user.SecretAccessKey, ec2.PublicIp, project.Name, goServerIp, ec2.KeyNumber))
	shouldReturn, result, err := f.addInboundRuleInBoGocd(&ec2.PublicIp)
	if shouldReturn {
		return result, err
	}

	return nil, nil
}

func (f *Ec2UserProjectFacade) InstallGoServer(input *useCaseDto.InstallCommand) (interface{}, error) {
	ec2, _ := f.ec2Svc.FindByInstanceId(&input.InstanceId)
	project, _ := f.projectSvc.Read(ec2.ProjectId)
	user, err := f.userSvc.FindByProjectIdAndKey(ec2.ProjectId, ec2.KeyNumber)

	if err != nil {
		return nil, err
	}
	ip := "0.0.0.0"
	f.ec2Svc.InstallGoServer(ec2DomainDto.InstallGocdCommandFrom(user.AccessKey, user.SecretAccessKey, ec2.PublicIp, project.Name, ec2.KeyNumber))
	keyName := "bohemiangocd0"
	f.ec2Svc.AddInboundRule(&user.AccessKey, &user.SecretAccessKey, &ip, &keyName)
	return nil, nil
}
