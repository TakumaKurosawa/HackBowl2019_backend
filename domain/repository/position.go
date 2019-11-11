package repository

import "backend_api/infrastructure/mysqlDB/model"

type PositionRepository interface {
	InsertPositionData(positionName string) error
	SelectPositionById(positionId int) (*model.Position, error)
	SelectPositionsByPositionIDs(positionIDs []int) ([]*model.Position, error)
}
