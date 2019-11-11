package mysqlDB

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

type ProjectUserRepositoryImpliment struct {
	DbConn *sql.DB
}

func NewProjectUserRepoImpl(conn *sql.DB) repository.ProjectUserRepository {
	return &ProjectUserRepositoryImpliment{
		DbConn: conn,
	}
}

func (repo *ProjectUserRepositoryImpliment) InsertProjectUserData(projectId, userId int) error {
	stmt, err := repo.DbConn.Prepare("INSERT INTO project_user(project_id, user_id) VALUES(?, ?)")
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	_, err = stmt.Exec(projectId, userId)
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}

	return nil
}

func (repo *ProjectUserRepositoryImpliment) SelectProjectUserByProjectId(projectId int) ([]*model.ProjectUser, error) {
	rows, err := repo.DbConn.Query("SELECT * FROM project_user WHERE project_id = ?", projectId)
	if err != nil {
		return nil, errors.Wrap(err, "DB Error(ProjectUser)")
	}

	defer rows.Close()
	return convertToProjectUser(rows)
}

// convertToProjectUser rowデータをProjectUserデータへ変換する
func convertToProjectUser(rows *sql.Rows) ([]*model.ProjectUser, error) {
	var results []*model.ProjectUser
	defer rows.Close()

	for rows.Next() {
		var projectId, UserId int

		if err := rows.Scan(&projectId, &UserId); err != nil {
			return nil, errors.Wrap(err, "DB Error(ProjectUser)")
		}

		row := model.ProjectUser{ProjectId: projectId, UserId: UserId}

		results = append(results, &row)
	}

	for i, _ := range results {
		log.Println(results[i])
	}

	return results, nil
}
