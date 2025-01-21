package user_handler

import (
	useCase "awsManager/api/user/cmd/application/useCase"
	useCaseDto "awsManager/api/user/cmd/application/useCase/dto/in"
	domain "awsManager/api/user/cmd/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	projectFcd useCase.IUserProjectFacade
	svc        domain.IService
}

func NewHandler(projectFcd useCase.IUserProjectFacade, svc domain.IService) *Handler {
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
	createdUser, err := h.projectFcd.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": createdUser})
}

func (h *Handler) FindInstanceOff(c *gin.Context) {
	projectName := c.Query("projectName")
	if projectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectName is required"})
		return
	}
	instanceOffUser, err := h.projectFcd.FindInstanceOff(projectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	keyNumbers := make([]uint, len(instanceOffUser))
	for i, user := range instanceOffUser {
		keyNumbers[i] = user.KeyNumber
	}
	c.JSON(http.StatusOK, gin.H{"userKeyNumbers": keyNumbers})
}
