package project_presentation

import (
	project_domain "awsManager/api/project/cmd/domain"
	project "awsManager/api/project/cmd/model"
	subProject_domain "awsManager/api/project/cmd/subProject/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	projectSvc    project_domain.IService
	subProjectSvc subProject_domain.IService
}

func NewHandler(projectSvc project_domain.IService, subProjectSvc subProject_domain.IService) *Handler {
	return &Handler{projectSvc: projectSvc, subProjectSvc: subProjectSvc}
}

func (h *Handler) Create(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}
	accountSuffix := c.Query("suffix")
	project, err := h.projectSvc.Create(name, accountSuffix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Duplicated project name"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project": project})
}

func (h *Handler) CreateSubProject(c *gin.Context) {

	var input struct {
		ProjectName    string `form:"projectName" binding:"required"`
		SubProjectName string `form:"subProjectName" binding:"required"`
		Group          string `form:"group" binding:"required"`
	}
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project, err := h.projectSvc.FindByName(input.ProjectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.subProjectSvc.Create(project.Id, input.SubProjectName, input.Group)
	subProjectNames := h.subProjectSvc.FindByProjectId(project.Id)
	c.JSON(http.StatusOK, gin.H{"project": project, "subProject": subProjectNames})
}

func (h *Handler) Read(c *gin.Context) {
	idStr := c.Query("id")
	id, shouldReturn := validateAndConvertId(idStr, c)
	if shouldReturn {
		return
	}

	project, err := h.projectSvc.Read(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to create project"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project": project})
}

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Query("id")
	id, shouldReturn := validateAndConvertId(idStr, c)
	if shouldReturn {
		return
	}

	var input struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProject := &project.Model{
		Id:   uint(id),
		Name: input.Name,
	}

	project, err := h.projectSvc.Update(updatedProject)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to create project"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project": project})
}

func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Query("id")
	name := c.Query("name")
	if idStr == "" && name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id or name are required"})
		return
	}
	if idStr == "" {
		err := h.projectSvc.DeleteByName(name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find project named, " + name})
		}
		return
	}
	id, shouldReturn := validateAndConvertId(idStr, c)
	if shouldReturn {
		return
	}
	project, err := h.projectSvc.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
	}
	c.JSON(http.StatusOK, gin.H{"deleted project": project})
}

func (h *Handler) List(c *gin.Context) {
	projects, err := h.projectSvc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func validateAndConvertId(idStr string, c *gin.Context) (uint, bool) {
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return 0, true
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return 0, true
	}
	if id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a non-negative integer"})
		return 0, true
	}
	return uint(id), false
}
