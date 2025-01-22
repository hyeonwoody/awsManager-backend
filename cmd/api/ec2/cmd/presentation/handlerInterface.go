package ec2_presentation

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Init(c *gin.Context)
}
