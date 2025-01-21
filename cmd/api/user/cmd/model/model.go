package user

import (
	"gorm.io/gorm"
)

type Model struct {
	ProjectId       uint   `gorm:"primaryKey"`
	KeyNumber       uint   `gorm:"primaryKey"`
	Password        string `gorm:"not null"`
	CanonicalUserId string `gorm:"not null"`
	AccessKey       string `gorm:"not null"`
	SecretAccessKey string `gorm:"not null"`
	SecurityGroupId string `gorm:"not null"`
	Ec2InstanceId   string
}

func (Model) TableName() string {
	return "user"
}

func (u *Model) BeforeSave(tx *gorm.DB) (err error) {
	if u.KeyNumber != 0 {
		return
	}
	var maxKeyNumber uint
	if err := tx.Model(&Model{}).
		Where("project_id = ?", u.ProjectId).
		Select("COALESCE(MAX(key_number), -1)").
		Scan(&maxKeyNumber).Error; err != nil {
		return err
	}
	u.KeyNumber = maxKeyNumber + 1
	return nil
}
