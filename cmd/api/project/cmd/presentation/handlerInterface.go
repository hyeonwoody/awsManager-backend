package project_presentation

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Create(c *gin.Context)
	CreateSubProject(c *gin.Context)
	Read(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
}
