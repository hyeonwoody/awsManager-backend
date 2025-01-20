package project_domain

import (
	project_infrastructure "awsManager/api/project/cmd/infrastructure"
	project "awsManager/api/project/cmd/model"
)

type Service struct {
	repo project_infrastructure.IRepository
}

func NewService(repo project_infrastructure.IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(name string, accountSuffix string) (*project.Model, error) {
	project := &project.Model{Name: name, AccountSuffix: accountSuffix}
	err := s.repo.Save(project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Service) Read(id uint) (*project.Model, error) {
	project, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Service) List() ([]project.Model, error) {
	return s.repo.FindAll()
}

func (s *Service) FindByName(name string) (*project.Model, error) {
	project, err := s.repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Service) Update(project *project.Model) (*project.Model, error) {
	return project, s.repo.Update(project)
}

func (s *Service) DeleteById(id uint) (*project.Model, error) {
	project, err := s.repo.DeleteById(id)
	return project, err
}

func (s *Service) DeleteByName(name string) error {
	return s.repo.DeleteByName(name)
}
