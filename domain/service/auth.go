package service

import (
	"backend_api/domain/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (svc *AuthService) Create(name string, email string, password string) (authToken string, err error) {
	authToken, err = createAuthToken()
	if err != nil {
		return "", err
	}
	hashedPassword, err := createHashedPassword(password)
	if err != nil {
		return "", err
	}

	err = svc.Repo.InsertUserData(name, email, hashedPassword, authToken)
	if err != nil {
		return "", err
	}

	return authToken, nil
}

func createHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func createAuthToken() (string, error) {
	tokenByte, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return tokenByte.String(), nil
}
