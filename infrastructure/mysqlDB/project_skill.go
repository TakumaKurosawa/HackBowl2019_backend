package mysqlDB

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"database/sql"
	"github.com/pkg/errors"
)

type ProjectSkillRepositoryImpliment struct {
	DbConn *sql.DB
}

func NewProjectSkillRepoImpl(conn *sql.DB) repository.ProjectSkillRepository {
	return &ProjectSkillRepositoryImpliment{
		DbConn: conn,
	}
}

func (repo *ProjectSkillRepositoryImpliment) InsertProjectSkillData(projectId, skillId int) error {
	stmt, err := repo.DbConn.Prepare("INSERT INTO project_skill(project_id, skill_id) VALUES(?, ?)")
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}
	_, err = stmt.Exec(projectId, skillId)
	if err != nil {
		return errors.Wrap(err, "DB Error")
	}

	return nil
}

func (repo *ProjectSkillRepositoryImpliment) SelectProjectSkillByProjectId(projectId int) ([]*model.ProjectSkill, error) {
	rows, err := repo.DbConn.Query("SELECT * FROM project_skill WHERE project_id = ?", projectId)
	if err != nil {
		return nil, errors.Wrap(err, "DB Error(ProjectSkill)")
	}

	defer rows.Close()
	return convertToProjectSkill(rows)
}

// convertToProjectSkill rowデータをProjectSkillデータへ変換する
func convertToProjectSkill(rows *sql.Rows) ([]*model.ProjectSkill, error) {
	var results []*model.ProjectSkill
	defer rows.Close()

	for rows.Next() {
		var projectId, skillId int

		if err := rows.Scan(&projectId, &skillId); err != nil {
			return nil, errors.Wrap(err, "DB Error(ProjectSkill)")
		}

		row := model.ProjectSkill{ProjectId: projectId, SkillId: skillId}

		results = append(results, &row)
	}

	return results, nil
}
