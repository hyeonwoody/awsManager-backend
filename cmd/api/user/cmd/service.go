package user

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindNextIndex(projectId uint) uint {
	return s.repo.FindNextIndex(projectId)
}

func (s *Service) Create(projectId, keyNubmber uint, password, accessKey, secretAccessKey string) (*Model, error) {
	if keyNubmber == 0 {
		keyNubmber = s.FindNextIndex(projectId)
	}
	user := &Model{ProjectId: projectId, KeyNumber: keyNubmber, Password: password, AccessKey: accessKey, SecretAccessKey: secretAccessKey}
	result := s.repo.Save(user)
	return user, result
}
