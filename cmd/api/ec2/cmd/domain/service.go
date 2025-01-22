package ec2_domain

import (
	ec2Businses "awsManager/api/ec2/cmd/business"
	businessDto "awsManager/api/ec2/cmd/business/dto"
	dto "awsManager/api/ec2/cmd/domain/dto"
	ec2Infrastructure "awsManager/api/ec2/cmd/infrastructure"
	ec2 "awsManager/api/ec2/cmd/model"
)

type Service struct {
	sdkBiz ec2Businses.IBusiness
	cliBiz ec2Businses.IBusiness
	repo   ec2Infrastructure.IRepository
}

func NewService(sdkBiz ec2Businses.IBusiness, cliBiz ec2Businses.IBusiness, repo ec2Infrastructure.IRepository) *Service {
	return &Service{sdkBiz: sdkBiz, cliBiz: cliBiz, repo: repo}
}

func (s *Service) DeleteExist(command *dto.DeleteCommand) error {
	s.sdkBiz.Delete(command)
	err := s.repo.DeleteByIdAndKeyNumber(command.ProjectId, command.KeyNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Create(command *dto.CreateCommand) (*ec2.Model, error) {
	ec2Instance, err := s.sdkBiz.Create(command)
	if err != nil {
		return nil, err
	}
	//s.Init(ec2Instance)
	ec2, err := s.repo.Save(businessDto.ModelFrom(command, ec2Instance))
	if err != nil {
		return nil, err
	}
	return ec2, nil
}

func (s *Service) Init(command *dto.InitWithPublicIpCommand) (*ec2.Model, error) {
	s.cliBiz.InitWithPublicIp(command)
	return nil, nil
}

func (s *Service) FindByInstanceId(instanceId *string) (*ec2.Model, error) {
	ec2, err := s.repo.FindByInstanceId(instanceId)
	if err != nil {
		return nil, err
	}
	return ec2, nil
}
