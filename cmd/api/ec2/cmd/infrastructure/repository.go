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
