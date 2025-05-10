package auth

import (
	"manajemen-user/internal/domain/user"
	"manajemen-user/utils"
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

	password, err := utils.HashedPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := user.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
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

	err = utils.CheckPasswordHash(user.Password, input.Password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID, user.Role.Name)
	if err != nil {
		return "", nil
	}

	return token, nil
}
