package subProject_infrastructure

import (
	subProject "awsManager/api/project/cmd/subProject/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(subProject *subProject.Model) (*subProject.Model, error) {
	result := r.db.Save(subProject)
	return subProject, result.Error
}

func (r *Repository) FindByProjectId(projectId uint) []subProject.Model {
	var subProjects []subProject.Model
	if err := r.db.Where("project_id = ?", projectId).
		Find(&subProjects).Error; err != nil {
		return nil
	}
	return subProjects

}
