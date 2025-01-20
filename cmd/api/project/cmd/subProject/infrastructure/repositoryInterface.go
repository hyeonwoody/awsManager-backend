package subProject_infrastructure

import subProject "awsManager/api/project/cmd/subProject/model"

type IRepository interface {
	Save(subProject *subProject.Model) (*subProject.Model, error)
	FindByProjectId(projectId uint) []subProject.Model
	// FindAll() ([]User, error)
	// Update(user *User) error
	// DeleteById(id uint) (*User, error)
}
