package project

import (
	"gorm.io/gorm"
)

type Model struct {
	Id            uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string `json:"name" gorm:"type:text;uniqueIndex:idx_project_name,length:30;not null"`
	AccountSuffix string `json:"accountSuffix" gorm:"not null"`
}

func (Model) TableName() string {
	return "project"
}

func (p *Model) BeforeCreate(tx *gorm.DB) error {
	var maxId uint
	if err := tx.Model(&Model{}).Select("COALESCE(MAX(id), 0)").Scan(&maxId).Error; err != nil {
		return err
	}
	p.Id = maxId + 1
	return nil
}
