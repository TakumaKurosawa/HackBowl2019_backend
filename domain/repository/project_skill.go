package repository

import "backend_api/infrastructure/mysqlDB/model"

type ProjectSkillRepository interface {
	InsertProjectSkillData(projectId, skillId int) error
	SelectProjectSkillByProjectId(projectId int) ([]*model.ProjectSkill, error)
}
