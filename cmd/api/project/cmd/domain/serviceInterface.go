package project_domain

import (
	project "awsManager/api/project/cmd/model"
)

type IService interface {
	Create(name string, accountSuffix string) (*project.Model, error)

	Read(id uint) (*project.Model, error)
	List() ([]project.Model, error)
	FindByName(name string) (*project.Model, error)

	Update(project *project.Model) (*project.Model, error)

	DeleteById(id uint) (*project.Model, error)
	DeleteByName(name string) error
}
