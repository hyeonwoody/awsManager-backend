package user

type IService interface {
	FindNextIndex(projectId uint) uint
	Create(projectId, keyNubmber uint, password, accessKey, secretAccessKey string) (*Model, error)
}
