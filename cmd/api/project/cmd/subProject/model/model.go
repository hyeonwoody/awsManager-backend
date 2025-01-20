package subProject

type Model struct {
	ProjectId uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"primaryKey;type:varchar(30);not null"`
	Group     string `json:"group" gorm:"type:varchar(30);not null"`
}

func (Model) TableName() string {
	return "sub_project"
}
