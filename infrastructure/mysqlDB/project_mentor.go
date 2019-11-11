package mysqlDB

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"database/sql"
	"github.com/pkg/errors"
)

type ProjectMentorRepositoryImpliment struct {
	DbConn *sql.DB
}

func NewProjectMentorRepoImpl(conn *sql.DB) repository.ProjectMentorRepository {
	return &ProjectMentorRepositoryImpliment{
		DbConn: conn,
	}
}

func (repo *ProjectMentorRepositoryImpliment) InsertProjectMentorData(projectId, mentorId int) error {
	stmt, err := repo.DbConn.Prepare("INSERT INTO project_mentor(project_id, mentor_id) VALUES(?, ?)")
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	_, err = stmt.Exec(projectId, mentorId)
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}

	return nil
}

func (repo *ProjectMentorRepositoryImpliment) SelectProjectMentorByProjectId(projectId int) (*model.ProjectMentor, error) {
	row := repo.DbConn.QueryRow("SELECT * FROM project_mentor WHERE project_id = ?", projectId)
	return convertToProjectMentor(row)
}

// convertToProjectMentor rowデータをProjectMentorデータへ変換する
func convertToProjectMentor(row *sql.Row) (*model.ProjectMentor, error) {
	projectMentor := model.ProjectMentor{}
	err := row.Scan(&projectMentor.ProjectId, &projectMentor.MentorId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "DB Error(ProjectMentor)")
		}
		return nil, errors.Wrap(err, "DB Error(ProjectMentor)")
	}
	return &projectMentor, nil
}
