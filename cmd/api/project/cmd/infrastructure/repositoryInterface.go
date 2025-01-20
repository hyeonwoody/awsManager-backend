package project_infrastructure

import (
	project "awsManager/api/project/cmd/model"
)

type IRepository interface {
	Save(project *project.Model) error
	FindById(id uint) (*project.Model, error)
	FindByName(name string) (*project.Model, error)
	FindAll() ([]project.Model, error)
	Update(project *project.Model) error
	DeleteById(id uint) (*project.Model, error)
	DeleteByName(name string) error
}
