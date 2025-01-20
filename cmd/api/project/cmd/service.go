package project

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(name string, accountSuffix string) (*Model, error) {
	project := &Model{Name: name, AccountSuffix: accountSuffix}
	err := s.repo.Save(project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Service) Read(id uint) (*Model, error) {
	project, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Service) List() ([]Model, error) {
	return s.repo.FindAll()
}

func (s *Service) FindByName(name string) (*Model, error) {
	project, err := s.repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Service) Update(project *Model) (*Model, error) {
	return project, s.repo.Update(project)
}

func (s *Service) DeleteById(id uint) (*Model, error) {
	project, err := s.repo.DeleteById(id)
	return project, err
}

func (s *Service) DeleteByName(name string) error {
	return s.repo.DeleteByName(name)
}
