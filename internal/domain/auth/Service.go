package auth

import (
	"manajemen-user/internal/domain/user"
	"manajemen-user/utils"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	ServiceRegister(input RegisterRequest) (*user.User, error)
	ServiceLogin(input LoginRequest) (string, error)
}

type service struct {
	Repo user.Repository
}

func NewService(repo user.Repository) Service {
	return &service{Repo: repo}
}

func (s *service) ServiceRegister(input RegisterRequest) (*user.User, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := user.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(password),
	}

	err = s.Repo.CreateUsers(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) ServiceLogin(input LoginRequest) (string, error) {
	user, err := s.Repo.FindByEmailWithRole(input.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID, user.Role.Name)
	if err != nil {
		return "", nil
	}

	return token, nil
}
