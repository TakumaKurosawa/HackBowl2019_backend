package entity

import "backend_api/infrastructure/mysqlDB/model"

type ProjectEntity struct {
	ProjectData *model.Project    `json:"project_data"`
	Position    []*PositionEntity `json:"position"`
	Skill       *SkillEntity      `json:"skill"`
	Mentor      *MentorEntity     `json:"mentor"`
	Users       []*UserEntity     `json:"users"`
}
