package ec2_presentation

import (
	useCase "awsManager/api/ec2/cmd/application/useCase"
	useCaseDto "awsManager/api/ec2/cmd/application/useCase/dto/in"
	domain "awsManager/api/ec2/cmd/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ec2Fcd useCase.IEc2UserProjectFacade
	svc    domain.IService
}

func NewHandler(ec2Fcd useCase.IEc2UserProjectFacade, svc domain.IService) *Handler {
	return &Handler{ec2Fcd: ec2Fcd, svc: svc}
}

func (h *Handler) Create(c *gin.Context) {
	var input *useCaseDto.CreateEc2Command
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ec2, err := h.ec2Fcd.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ec2": ec2})
}

func (h *Handler) Init(c *gin.Context) {
	var input *useCaseDto.InitEc2Command
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ec2, err := h.ec2Fcd.AddMemory(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ec2": ec2})
}

func (h *Handler) Attach(c *gin.Context) {
	var input *useCaseDto.AttachEbsVolumeCommand
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.ec2Fcd.AttachEbsVolume(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"attach": nil})
}

func (h *Handler) InstallDocker(c *gin.Context) {
	var input *useCaseDto.InstallCommand
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ec2, err := h.ec2Fcd.InstallDocker(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ec2": ec2})
}

func (h *Handler) InstallDockerNginx(c *gin.Context) {
	var input *useCaseDto.InstallCommand
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ec2, err := h.ec2Fcd.InstallDockerNginx(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ec2": ec2})
}

func (h *Handler) InstallGoAgent(c *gin.Context) {
	var input *useCaseDto.InstallCommand
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ec2, err := h.ec2Fcd.InstallGoAgent(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ec2": ec2})
}

func (h *Handler) InstallDockerGoAgent(c *gin.Context) {
	var input *useCaseDto.InstallCommand
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ec2, err := h.ec2Fcd.InstallDockerGoAgent(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ec2": ec2})
}

func (h *Handler) InstallGoServer(c *gin.Context) {
	var input *useCaseDto.InstallCommand
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ec2, err := h.ec2Fcd.InstallGoServer(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ec2": ec2})
}
