package project

import (
	di "awsManager/dependencyInjection"

	"github.com/gin-gonic/gin"
)

func Main(container *di.Container, router *gin.Engine) {
	projectGroup := router.Group("/projects")
	{
		projectGroup.POST("", container.ProjectHandler.Create)
		projectGroup.GET("/list", container.ProjectHandler.List)
		projectGroup.DELETE("", container.ProjectHandler.Delete)

		projectGroup.POST("/sub-projects", container.ProjectHandler.CreateSubProject)
	}
}
