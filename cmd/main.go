package main

import (
	project "awsManager/api/project"
	"awsManager/api/user"
	"awsManager/database"
	di "awsManager/dependencyInjection"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDatabase()
	database.Migrate()

	container := di.Container{}
	container.Init(database.DB)

	router := gin.Default()

	user.Main(&container, router)
	project.Main(&container, router)

	router.Run(":10000")
}
