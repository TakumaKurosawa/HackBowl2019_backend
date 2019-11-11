package repository

import "backend_api/infrastructure/mysqlDB/model"

type MentorRepository interface {
	InsertMentorData(name, email, password string) error
	SelectMentorById(mentorId int) (*model.Mentor, error)
	SelectMentorsByMentorIDs(mentorIDs []int) ([]*model.Mentor, error)
}
