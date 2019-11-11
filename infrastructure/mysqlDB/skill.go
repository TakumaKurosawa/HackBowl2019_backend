package mysqlDB

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type SkillRepositoryImpliment struct {
	DbConn *sql.DB
}

func NewSkillRepoImpl(conn *sql.DB) repository.SkillRepository {
	return &SkillRepositoryImpliment{
		DbConn: conn,
	}
}

func (repo *SkillRepositoryImpliment) InsertSkillData(skillName string) error {
	stmt, err := repo.DbConn.Prepare("INSERT INTO skill(skill_name) VALUES(?)")
	if err != nil {
		return errors.Wrap(err, "read")
	}
	_, err = stmt.Exec(skillName)
	if err != nil {
		return errors.Wrap(err, "read")
	}

	return nil
}

func (repo *SkillRepositoryImpliment) SelectSkillById(skillId int) (*model.Skill, error) {
	row := repo.DbConn.QueryRow("SELECT * FROM skill WHERE id = ?", skillId)
	return convertToSkill(row)
}

func (repo *SkillRepositoryImpliment) SelectSkillsBySkillIDs(skillIDs []int) ([]*model.Skill, error) {
	var queryStr string

	for i, _ := range skillIDs {
		if i+1 == len(skillIDs) {
			queryStr += fmt.Sprintf("%v", skillIDs[i])
			break
		}
		queryStr += fmt.Sprintf("%v,", skillIDs[i])
	}

	rows, err := repo.DbConn.Query(fmt.Sprintf("SELECT * FROM skill WHERE id IN (%v)", queryStr))
	if err != nil {
		return nil, errors.Wrap(err, "DB Error")
	}

	defer rows.Close()
	return convertToSkills(rows)
}

// convertToProject rowデータをProjectデータへ変換する
func convertToSkill(row *sql.Row) (*model.Skill, error) {
	skill := model.Skill{}
	err := row.Scan(&skill.Id, &skill.SkillName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "DB Error")
		}
		return nil, errors.Wrap(err, "DB Error")
	}
	return &skill, nil
}

func convertToSkills(rows *sql.Rows) ([]*model.Skill, error) {
	var results []*model.Skill
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return nil, errors.Wrap(err, "DB Error")
		}

		row := model.Skill{Id: id, SkillName: name}

		results = append(results, &row)
	}

	return results, nil
}
