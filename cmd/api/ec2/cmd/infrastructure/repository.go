package ec2_infrastructure

import (
	ec2 "awsManager/api/ec2/cmd/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ec2 *ec2.Model) (*ec2.Model, error) {
	result := r.db.Save(ec2)
	return ec2, result.Error
}

func (r *Repository) DeleteByIdAndKeyNumber(projectId, keyNumber uint) error {
	result := r.db.Where("project_id = ? AND key_number = ?", projectId, keyNumber).Delete(&ec2.Model{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) FindByInstanceId(instanceId *string) (*ec2.Model, error) {
	var ec2 ec2.Model
	result := r.db.Table("ec2").Where("instance_id = ?", instanceId).
		First(&ec2)
	return &ec2, result.Error
}

func (r *Repository) FindByProjectIdAndKey(projectId, keyNumber uint) (*ec2.Model, error) {
	var ec2 ec2.Model
	result := r.db.Table("ec2").Where("project_id = ? AND key_number = ?", projectId, keyNumber).
		First(&ec2)
	return &ec2, result.Error
}
