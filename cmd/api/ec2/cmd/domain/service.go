package ec2_domain

import (
	ec2Businses "awsManager/api/ec2/cmd/business"
	dto "awsManager/api/ec2/cmd/dto"
	ec2Infrastructure "awsManager/api/ec2/cmd/infrastructure"
	ec2 "awsManager/api/ec2/cmd/model"
)

type Service struct {
	biz  ec2Businses.IBusiness
	repo ec2Infrastructure.IRepository
}

func NewService(biz ec2Businses.IBusiness, repo ec2Infrastructure.IRepository) *Service {
	return &Service{biz: biz, repo: repo}
}

func (s *Service) Create(command *dto.CreateCommand) (*ec2.Model, error) {
	ec2Instance, err := s.biz.Create(command)
	if err != nil {
		return nil, err
	}
	//s.Init(ec2Instance)
	ec2, err := s.repo.Save(dto.ModelFrom(command, ec2Instance))
	if err != nil {
		return nil, err
	}
	return ec2, nil
}

func (s *Service) Init(ec2Instance *dto.Ec2Instance) {
	s.biz.Init(ec2Instance)
}
