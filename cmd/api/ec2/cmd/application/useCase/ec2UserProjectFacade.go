package ec2_useCase

import (
	useCasedto "awsManager/api/ec2/cmd/application/useCase/dto/in"
	ec2Domain "awsManager/api/ec2/cmd/domain"
	dto "awsManager/api/ec2/cmd/dto"
	ec2 "awsManager/api/ec2/cmd/model"
	projectDomain "awsManager/api/project/cmd/domain"
	userDomain "awsManager/api/user/cmd/domain"
	"fmt"
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

func (f *Ec2UserProjectFacade) Init(command *useCasedto.InitEc2Command) (*ec2.Model, error) {
	project, err := f.projectSvc.FindByName(command.ProjectName)
	if err != nil {
		return nil, err
	}
	user, err := f.userSvc.FindByProjectIdAndKey(project.Id, command.KeyNumber)
	if err != nil {
		return nil, err
	}
	if user.Ec2InstanceId != "" {
		return nil, fmt.Errorf("user already have instance")
	}

	ec2, err := f.ec2Svc.Create(dto.CreateCommandFrom(project.Name, command.Ami, command.InstanceType, user.AccessKey, user.SecretAccessKey, project.Id, user.KeyNumber))
	if err != nil {
		return nil, err
	}
	user.Ec2InstanceId = ec2.InstanceId
	go f.userSvc.Save(user)
	return ec2, nil
}
