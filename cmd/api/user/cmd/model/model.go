package user

import (
	"time"

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
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}

func (Model) TableName() string {
	return "user"
}

func (u *Model) BeforeSave(tx *gorm.DB) (err error) {
	if u.KeyNumber == 0 {
		return nil
	}
	if u.KeyNumber != 0 {
		return nil
	}
	var maxKeyNumber int
	if err := tx.Model(&Model{}).
		Where("project_id = ?", u.ProjectId).
		Select("COALESCE(MAX(key_number), 0)").
		Scan(&maxKeyNumber).Error; err != nil {
		return err
	}
	u.KeyNumber = uint(maxKeyNumber) + 1
	return nil
}
