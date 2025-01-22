package domain

import (
	Business "awsManager/api/user/cmd/business"
	Infrastructure "awsManager/api/user/cmd/infrastructure"
	user "awsManager/api/user/cmd/model"
	"fmt"
)

type Service struct {
	biz  Business.IBusiness
	repo Infrastructure.IRepository
}

func NewService(biz Business.IBusiness, repo Infrastructure.IRepository) *Service {
	return &Service{biz: biz, repo: repo}
}

func (s *Service) FindNextIndex(projectId uint) uint {
	return s.repo.FindNextIndex(projectId)
}

func (s *Service) Create(projectId, keyNubmber uint, projectName, password, accessKey, secretAccessKey string) (*user.Model, error) {
	if keyNubmber == 0 {
		keyNubmber = s.FindNextIndex(projectId)
	}

	s.biz.CreateCredentialConfigure(projectName, accessKey, secretAccessKey, keyNubmber)

	user := &user.Model{ProjectId: projectId, KeyNumber: keyNubmber, Password: password, AccessKey: accessKey, SecretAccessKey: secretAccessKey}
	result := s.repo.Save(user)
	return user, result
}

func (s *Service) FindByProjectIdAndKey(id uint, keyNumber uint) (*user.Model, error) {
	user, err := s.repo.FindByProjectIdAndKey(id, keyNumber)
	if user == nil {
		return user, fmt.Errorf("failed to find user: %w", err)
	}
	return user, err
}

func (s *Service) FindInstanceOff(projectId uint) ([]user.Model, error) {
	user, err := s.repo.FindInstanceOff(projectId)
	if user == nil {
		return user, fmt.Errorf("failed to find user: %w", err)
	}
	return user, err
}

func (s *Service) Save(user *user.Model) error {
	s.repo.Save(user)
	return nil
}
