package responseBody

import (
	"backend_api/domain/entity"
	"database/sql"
)

type Project struct {
	ProjectId   int                      `json:"project_id"`
	ProjectName string                   `json:"project_name"`
	StartDate   sql.NullString           `json:"start_date"`
	Level       int                      `json:"level"`
	Position    []*entity.PositionEntity `json:"position"`
	Skill       *entity.SkillEntity      `json:"skill"`
	Mentor      *entity.MentorEntity     `json:"mentor"`
	Users       []*entity.UserEntity     `json:"users"`
}
