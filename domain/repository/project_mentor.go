package repository

import "backend_api/infrastructure/mysqlDB/model"

type ProjectMentorRepository interface {
	InsertProjectMentorData(projectId, mentorId int) error
	SelectProjectMentorByProjectId(projectId int) (*model.ProjectMentor, error)
}
