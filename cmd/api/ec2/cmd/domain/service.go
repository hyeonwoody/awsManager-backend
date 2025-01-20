package ec2_domain

import (
	//ec2 "awsManager/api/ec2/cmd/model"
	ec2_infrastructure "awsManager/api/ec2/cmd/infrastructure"
)

type Service struct {
	repo ec2_infrastructure.IRepository
}

func NewService(repo ec2_infrastructure.IRepository) *Service {
	return &Service{repo: repo}
}
