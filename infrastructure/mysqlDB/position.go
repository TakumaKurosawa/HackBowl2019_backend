package mysqlDB

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type PositionRepositoryImpliment struct {
	DbConn *sql.DB
}

func NewPositionRepoImpl(conn *sql.DB) repository.PositionRepository {
	return &PositionRepositoryImpliment{
		DbConn: conn,
	}
}

func (repo *PositionRepositoryImpliment) InsertPositionData(positionName string) error {
	stmt, err := repo.DbConn.Prepare("INSERT INTO position(position_name) VALUES(?)")
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	_, err = stmt.Exec(positionName)
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}

	return nil
}

func (repo *PositionRepositoryImpliment) SelectPositionById(positionId int) (*model.Position, error) {
	row := repo.DbConn.QueryRow("SELECT * FROM position WHERE id = ?", positionId)
	return convertToPosition(row)
}

func (repo *PositionRepositoryImpliment) SelectPositionsByPositionIDs(positionIDs []int) ([]*model.Position, error) {
	var queryStr string

	for i, _ := range positionIDs {
		if i+1 == len(positionIDs) {
			queryStr += fmt.Sprintf("%v", positionIDs[i])
			break
		}
		queryStr += fmt.Sprintf("%v,", positionIDs[i])
	}

	rows, err := repo.DbConn.Query(fmt.Sprintf("SELECT * FROM position WHERE id IN (%v)", queryStr))
	if err != nil {
		return nil, errors.Wrap(err, "DB Error")
	}

	defer rows.Close()
	return convertToPositions(rows)
}

// convertToPosition rowデータをPositionデータへ変換する
func convertToPosition(row *sql.Row) (*model.Position, error) {
	position := model.Position{}
	err := row.Scan(&position.Id, &position.PositionName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "DB Error")
		}
		return nil, errors.Wrap(err, "DB Error")
	}
	return &position, nil
}

func convertToPositions(rows *sql.Rows) ([]*model.Position, error) {
	var results []*model.Position
	defer rows.Close()

	for rows.Next() {
		var id int
		var positionName string

		if err := rows.Scan(&id, &positionName); err != nil {
			return nil, errors.Wrap(err, "DB Error")
		}

		row := model.Position{Id: id, PositionName: positionName}

		results = append(results, &row)
	}

	return results, nil
}
