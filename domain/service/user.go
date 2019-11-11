package service

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (svc *UserService) GetUserName(userId string) (name string, err error) {
	selectedUser, err := svc.Repo.SelectUserByUserId(userId)
	if err != nil {
		return "", err
	}

	return selectedUser.Name, err
}

func (svc *UserService) ChangeUserName(token, name string) error {
	err := svc.Repo.UpdateUserName(token, name)
	if err != nil {
		return err
	}
	return nil
}

func (svc *UserService) ChangeUserEmail(token, email string) error {
	err := svc.Repo.UpdateUserEmail(token, email)
	if err != nil {
		return err
	}
	return nil
}

func (svc *UserService) ChangeUserData(token, name, email string) error {
	err := svc.Repo.UpdateUserData(token, name, email)
	if err != nil {
		return err
	}
	return nil
}

func (svc *UserService) GetUserByAuthToken(token string) (*model.User, error) {
	result, err := svc.Repo.SelectByAuthToken(token)
	if err != nil {
		return nil, err
	}
	return result, nil
}
