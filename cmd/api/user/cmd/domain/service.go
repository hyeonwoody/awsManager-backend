package domain

import (
	user_infrastructure "awsManager/api/user/cmd/infrastructure"
	user "awsManager/api/user/cmd/model"
)

type Service struct {
	repo user_infrastructure.IRepository
}

func NewService(repo user_infrastructure.IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindNextIndex(projectId uint) uint {
	return s.repo.FindNextIndex(projectId)
}

func (s *Service) Create(projectId, keyNubmber uint, password, accessKey, secretAccessKey string) (*user.Model, error) {
	if keyNubmber == 0 {
		keyNubmber = s.FindNextIndex(projectId)
	}
	user := &user.Model{ProjectId: projectId, KeyNumber: keyNubmber, Password: password, AccessKey: accessKey, SecretAccessKey: secretAccessKey}
	result := s.repo.Save(user)
	return user, result
}
