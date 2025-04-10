package user

import (
	di "awsManager/dependencyInjection"

	"github.com/gin-gonic/gin"
)

func Main(container *di.Container, router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		// userRepository := user.NewRepository(db)
		// userService := user.NewService(userRepository)
		// userHandler := user.NewHandler(userService)
		userGroup.GET("all", container.UserHandler.ReadAll)
		userGroup.GET("/next-index", container.UserHandler.FindNextIndex)
		userGroup.POST("", container.UserHandler.Create)
		userGroup.GET("/instance-off", container.UserHandler.FindInstanceOff)
	}
}
