package repository

import "backend_api/infrastructure/mysqlDB/model"

type ProjectPositionRepository interface {
	InsertProjectPositionData(projectId, positionId, limitNum int) error
	SelectProjectPositionByProjectId(projectId int) ([]*model.ProjectPosition, error)
}
