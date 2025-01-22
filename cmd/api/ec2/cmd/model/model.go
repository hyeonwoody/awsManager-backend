package model

import "time"

type InstanceType int

const (
	T2Micro InstanceType = iota
)

type Model struct {
	InstanceId string    `gorm:"primaryKey;not null"`
	ProjectId  uint      `gorm:"not null"`
	KeyNumber  uint      `gorm:"not null"`
	Ami        string    `gorm:"not null"`
	PublicIp   string    `gorm:"not null"`
	PrivateIp  string    `gorm:"not null"`
	CicdOn     bool      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (Model) TableName() string {
	return "ec2"
}
