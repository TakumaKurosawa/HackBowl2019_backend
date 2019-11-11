package mysqlDB

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"database/sql"
	"github.com/pkg/errors"
)

type ProjectPositionRepositoryImpliment struct {
	DbConn *sql.DB
}

func NewProjectPositionRepoImpl(conn *sql.DB) repository.ProjectPositionRepository {
	return &ProjectPositionRepositoryImpliment{
		DbConn: conn,
	}
}

func (repo *ProjectPositionRepositoryImpliment) InsertProjectPositionData(projectId, positionId, limitNum int) error {
	stmt, err := repo.DbConn.Prepare("INSERT INTO project_position(project_id, position_id, limit_num) VALUES(?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "DB Error(ProjectPosition)")
	}
	_, err = stmt.Exec(projectId, positionId, limitNum)
	if err != nil {
		return errors.Wrap(err, "DB Error(ProjectPosition)")
	}

	return nil
}

func (repo *ProjectPositionRepositoryImpliment) SelectProjectPositionByProjectId(projectId int) ([]*model.ProjectPosition, error) {
	rows, err := repo.DbConn.Query("SELECT * FROM project_position WHERE project_id = ?", projectId)
	if err != nil {
		return nil, errors.Wrap(err, "DB Error(ProjectPosition)")
	}

	defer rows.Close()
	return convertToProjectPosition(rows)
}

// convertToProjectPosition rowデータをProjectPositionデータへ変換する
func convertToProjectPosition(rows *sql.Rows) ([]*model.ProjectPosition, error) {
	var results []*model.ProjectPosition
	defer rows.Close()

	for rows.Next() {
		var projectId, positionId, limitNum int

		if err := rows.Scan(&projectId, &positionId, &limitNum); err != nil {
			return nil, errors.Wrap(err, "DB Error(ProjectPosition)")
		}

		row := model.ProjectPosition{ProjectId: projectId, PositionId: positionId, LimitNum: limitNum}

		results = append(results, &row)
	}

	return results, nil
}
