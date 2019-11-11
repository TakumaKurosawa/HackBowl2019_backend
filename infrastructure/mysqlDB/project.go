package mysqlDB

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"database/sql"
	"github.com/pkg/errors"
)

type ProjectRepositoryImpliment struct {
	DbConn *sql.DB
}

func NewProjectRepoImpl(conn *sql.DB) repository.ProjectRepository {
	return &ProjectRepositoryImpliment{
		DbConn: conn,
	}
}

func (repo *ProjectRepositoryImpliment) SelectProjectByProjectId(projectId int) (*model.Project, error) {
	row := repo.DbConn.QueryRow("SELECT * FROM project WHERE id = ?", projectId)
	return convertToProject(row)
}

func (repo *ProjectRepositoryImpliment) SelectAllProjects() ([]*model.Project, error) {
	rows, err := repo.DbConn.Query("SELECT * FROM project")
	if err != nil {
		return nil, errors.Wrap(err, "DB Error")
	}

	defer rows.Close()
	return convertToProjects(rows)
}

func convertToProjects(rows *sql.Rows) ([]*model.Project, error) {
	var results []*model.Project
	defer rows.Close()

	for rows.Next() {
		var id, level int
		var name string
		var startDate sql.NullString

		if err := rows.Scan(&id, &name, &level, &startDate); err != nil {
			return nil, errors.Wrap(err, "DB Error")
		}

		row := model.Project{Id: id, Name: name, Level: level, StartDate: startDate}

		results = append(results, &row)
	}

	return results, nil
}

// convertToProject rowデータをProjectデータへ変換する
func convertToProject(row *sql.Row) (*model.Project, error) {
	project := model.Project{}
	err := row.Scan(&project.Id, &project.Name, &project.Level, &project.StartDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "DB Error")
		}
		return nil, errors.Wrap(err, "DB Error")
	}
	return &project, nil
}
