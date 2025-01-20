package user_handler

import "github.com/gin-gonic/gin"

type IHandler interface {
	FindNextIndex(c *gin.Context)
	Create(c *gin.Context)
}
