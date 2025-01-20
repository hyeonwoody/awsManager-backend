package database

import (
	project "awsManager/api/project/cmd/model"
	subProject "awsManager/api/project/cmd/subProject/model"
	user "awsManager/api/user/cmd/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3307)/awsManager?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}

// Migrate the schema
func Migrate() {
	err := DB.AutoMigrate(
		&project.Model{},
		&user.Model{},
		&subProject.Model{},
	)
	if err != nil {
		panic("Failed to migrate database : " + err.Error())
	}
}
