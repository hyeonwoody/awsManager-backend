package user

type IRepository interface {
	FindNextIndex(projectId uint) uint
	Save(user *Model) error
	// Save(user *User) error
	// FindById(id uint) (*User, error)
	// FindAll() ([]User, error)
	// Update(user *User) error
	// DeleteById(id uint) (*User, error)
}
