package project

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(project *Model) error {
	result := r.db.Save(project)
	return result.Error
}

func (r *Repository) FindById(id uint) (*Model, error) {
	var project Model
	result := r.db.First(&project, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func (r *Repository) FindByName(name string) (*Model, error) {
	var project *Model
	result := r.db.Where("name = ?", name).First(&project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *Repository) FindAll() ([]Model, error) {
	var projects []Model
	if err := r.db.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *Repository) Update(project *Model) error {
	return r.db.Save(project).Error
}

func (r *Repository) DeleteById(id uint) (*Model, error) {
	var project Model
	if err := r.db.First(&project, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Delete(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *Repository) DeleteByName(name string) error {
	result := r.db.Where("name = ?", name).Delete(&Model{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
