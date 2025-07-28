package auth

import (
	"errors"
	"log"
	"manajemen-user/internal/domain/user"
	appErrors "manajemen-user/internal/errors"
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
		log.Printf("Failed to hash password for user %s: %v", input.Email, err)
		return nil, appErrors.ErrInternal
	}

	user := user.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
	}

	err = s.Repo.CreateUsers(&user)
	if err != nil {
		log.Printf("Failed to create user %s in database: %v", user.Email, err)
		return nil, appErrors.ErrInternal
	}

	return &user, nil
}

func (s *service) ServiceLogin(input LoginRequest) (string, error) {
	user, err := s.Repo.FindByEmailWithRole(input.Email)

	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			log.Printf("Login attempt failed: User with email %s not found", input.Email)
			return "", appErrors.ErrNotFound
		}

		log.Printf("Failed to query user with email %s: %v", input.Email, err)
		return "", appErrors.ErrInternal
	}

	err = utils.CheckPasswordHash(user.Password, input.Password)
	if err != nil {
		log.Printf("Invalid password attempt for user %s", user.Email)
		return "", appErrors.ErrInternal
	}

	token, err := utils.GenerateToken(user.ID, user.Role.Name)
	if err != nil {
		log.Printf("Failed to generate JWT for user ID %d: %v", user.ID, err)
		return "", appErrors.ErrInternal
	}

	return token, nil
}
