package model

type InstanceType int

const (
	T2Micro InstanceType = iota
)

type Model struct {
	ProjectId uint `gorm:"primaryKey;not null"`
	KeyNumber uint `gorm:"primaryKey;not null"`
}

func (it *InstanceType) Int() int {
	return int(*it)
}
