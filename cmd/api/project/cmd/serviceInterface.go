package project

type IService interface {
	Create(name string, accountSuffix string) (*Model, error)

	Read(id uint) (*Model, error)
	List() ([]Model, error)
	FindByName(name string) (*Model, error)

	Update(project *Model) (*Model, error)

	DeleteById(id uint) (*Model, error)
	DeleteByName(name string) error
}
