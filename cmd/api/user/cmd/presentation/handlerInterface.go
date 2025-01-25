package user_handler

import "github.com/gin-gonic/gin"

type IHandler interface {
	ReadAll(c *gin.Context)
	FindNextIndex(c *gin.Context)
	Create(c *gin.Context)
	FindInstanceOff(c *gin.Context)
}
