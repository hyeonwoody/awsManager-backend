package user_infrastructure

import (
	user "awsManager/api/user/cmd/model"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(user *user.Model) error {
	result := r.db.Save(user)
	return result.Error
}

func (r *Repository) FindById(id uint) (*user.Model, error) {
	var user user.Model
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) FindNextIndex(projectId uint) uint {
	var nextIndex int
	result := r.db.Table("user").Where("project_id = ?", projectId).
		Select("COALESCE(MAX(key_number), 0)").
		Scan(&nextIndex)
	if result.Error != nil {
		return 0
	}
	if nextIndex == 0 {
		return 0
	}
	return uint(nextIndex) + 1
}

func (r *Repository) FindByProjectIdAndKey(projectId uint, keyNumber uint) (*user.Model, error) {
	var user user.Model
	result := r.db.Table("user").Where("project_id = ? AND key_number = ?", projectId, keyNumber).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) ReadAll() ([]user.Model, error) {
	var users []user.Model
	if result := r.db.Table("user").Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *Repository) FindInstanceOff(projectId uint) ([]user.Model, error) {
	var users []user.Model

	if err := r.db.Table("user").Where("project_id = ? AND ec2_instance_id IS NULL", projectId).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) Update(user *user.Model) error {
	result := r.db.Table("user").Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
