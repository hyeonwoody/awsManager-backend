package ec2_presentation

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Create(c *gin.Context)
	Init(c *gin.Context)
	Attach(c *gin.Context)
	InstallDocker(c *gin.Context)
	InstallDockerNginx(c *gin.Context)
	InstallGoAgent(c *gin.Context)
	InstallDockerGoAgent(c *gin.Context)
	InstallGoServer(c *gin.Context)
}
