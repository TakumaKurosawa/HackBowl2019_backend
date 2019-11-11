package service

import (
	"backend_api/domain/repository"
	"backend_api/infrastructure/mysqlDB/model"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	Repo repository.UserRepository
}

func NewLoginService(repo repository.UserRepository) *LoginService {
	return &LoginService{
		Repo: repo,
	}
}

func (svc *LoginService) Authenticate(email string, password string) (selectedUser *model.User, err error) {
	selectedUserData, err := svc.Repo.SelectUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = passwordVerify(selectedUserData.Password, password)
	if err != nil {
		return nil, err
	}

	return selectedUserData, nil
}

func passwordVerify(storedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword))
}