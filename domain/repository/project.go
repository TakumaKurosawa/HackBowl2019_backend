package repository

import "backend_api/infrastructure/mysqlDB/model"

type ProjectRepository interface {
	SelectProjectByProjectId(projectId int) (*model.Project, error)
	SelectAllProjects() ([]*model.Project, error)
}
