package repository

import "backend_api/infrastructure/mysqlDB/model"

type ProjectUserRepository interface {
	InsertProjectUserData(projectId, userId int) error
	SelectProjectUserByProjectId(projectId int) ([]*model.ProjectUser, error)
}
