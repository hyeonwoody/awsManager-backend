package project

type IRepository interface {
	Save(project *Model) error
	FindById(id uint) (*Model, error)
	FindByName(name string) (*Model, error)
	FindAll() ([]Model, error)
	Update(project *Model) error
	DeleteById(id uint) (*Model, error)
	DeleteByName(name string) error
}
