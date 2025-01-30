package ec2

import (
	di "awsManager/dependencyInjection"

	"github.com/gin-gonic/gin"
)

func Main(container *di.Container, router *gin.Engine) {
	ec2Group := router.Group("/ec2")
	{
		ec2Group.POST("", container.Ec2Handler.Create)
		ec2Group.PATCH("attach", container.Ec2Handler.Attach)
		ec2Group.POST("swapfile", container.Ec2Handler.Init)
		ec2Group.POST("install/docker", container.Ec2Handler.InstallDocker)
		ec2Group.POST("install/docker-nginx", container.Ec2Handler.InstallDockerNginx)
		ec2Group.POST("install/go-agent", container.Ec2Handler.InstallGoAgent)
		ec2Group.POST("install/docker-go-agent", container.Ec2Handler.InstallDockerGoAgent)
		ec2Group.POST("install/go-server", container.Ec2Handler.InstallGoServer)
	}
}
