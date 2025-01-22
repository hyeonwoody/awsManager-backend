package ec2

import (
	di "awsManager/dependencyInjection"

	"github.com/gin-gonic/gin"
)

func Main(container *di.Container, router *gin.Engine) {
	ec2Group := router.Group("/ec2")
	{
		ec2Group.POST("", container.Ec2Handler.Create)
		ec2Group.PATCH("/init", container.Ec2Handler.Init)
	}
}
