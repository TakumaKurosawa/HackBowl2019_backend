package repository

import "backend_api/infrastructure/mysqlDB/model"

type UserRepository interface {
	InsertUserData(name string, email string, password string, authToken string) error
	SelectByAuthToken(authToken string) (*model.User, error)
	SelectUserByEmail(email string) (*model.User, error)
	SelectUserByUserId(userId string) (*model.User, error)
	SelectUsersByUserIDs(userId []int) ([]*model.User, error)
	UpdateUserName(token string, name string) error
	UpdateUserEmail(token string, email string) error
	UpdateUserData(token string, name string, email string) error
}
