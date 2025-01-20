package subProject_domain

import subProject "awsManager/api/project/cmd/subProject/model"

type IService interface {
	Create(projectId uint, name, group string) []*subProject.Model
	FindByProjectId(id uint) []string
}
