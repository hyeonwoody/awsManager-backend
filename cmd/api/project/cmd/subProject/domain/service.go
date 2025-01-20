package subProject_domain

import (
	subProject_infra "awsManager/api/project/cmd/subProject/infrastructure"
	subProject "awsManager/api/project/cmd/subProject/model"
	"strings"
)

type Service struct {
	repo subProject_infra.IRepository
}

func NewService(repo subProject_infra.IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(projectId uint, name, group string) []*subProject.Model {

	var savedSubProjects []*subProject.Model
	names := strings.Split(name, ",")
	for _, eachName := range names {
		subProject := subProject.Model{
			ProjectId: projectId,
			Name:      eachName,
			Group: func() string {
				if group == "" {
					return name
				}
				return group
			}(),
		}
		createdSubProject, err := s.repo.Save(&subProject)
		if err != nil {
			return nil
		}
		savedSubProjects = append(savedSubProjects, createdSubProject)
	}
	return savedSubProjects
}

func (s *Service) FindByProjectId(projectId uint) []string {
	subProjects := s.repo.FindByProjectId(projectId)

	var names []string
	for _, subProject := range subProjects {
		names = append(names, subProject.Name)
	}
	return names
}
