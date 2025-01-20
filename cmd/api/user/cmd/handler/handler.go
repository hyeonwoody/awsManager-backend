package user_handler

import (
	user "awsManager/api/user/cmd"
	useCase "awsManager/api/user/cmd/useCase"
	useCaseDto "awsManager/api/user/cmd/useCase/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "net/http"

//"fmt"

type Handler struct {
	projectFcd useCase.IUserProjectFacade
	svc        user.IService
}

func NewHandler(projectFcd useCase.IUserProjectFacade, svc user.IService) *Handler {
	return &Handler{
		projectFcd: projectFcd,
		svc:        svc}
}

func (h *Handler) FindNextIndex(c *gin.Context) {
	projectName := c.Query("projectName")
	if projectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectName is required"})
		return
	}
	nextIndex, err := h.projectFcd.FindNextIndex(projectName)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"nextIndex": nextIndex})
}

func (h *Handler) Create(c *gin.Context) {
	projectName := c.Query("projectName")
	if projectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectName is required"})
		return
	}
	var input useCaseDto.CreateUserCommand
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.projectFcd.CreateUser(input)
}
