package user

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(user *Model) error {
	result := r.db.Save(user)
	return result.Error
}

func (r *Repository) FindById(id uint) (*Model, error) {
	var user Model
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) FindNextIndex(projectId uint) uint {
	var nextIndex uint
	result := r.db.Table("user").Where("project_id = ?", projectId).
		Select("COALESCE(MAX(key_number), 0)").
		Scan(&nextIndex)
	if result.Error != nil {
		return 0
	}
	return nextIndex
}
