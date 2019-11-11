package repository

import "backend_api/infrastructure/mysqlDB/model"

type SkillRepository interface {
	InsertSkillData(skillName string) error
	SelectSkillById(skillId int) (*model.Skill, error)
	SelectSkillsBySkillIDs(skillIDs []int) ([]*model.Skill, error)
}
