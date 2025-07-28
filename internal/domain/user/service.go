package user

import (
	"errors"
	"fmt"
	"log"
	appErrors "manajemen-user/internal/errors"
	"manajemen-user/utils"
)

type Service interface {
	ServiceGetUsers() ([]User, error)
	ServiceGetUsersByID(id string) (*User, error)
	ServiceCreateUsers(input CreateUsersRequest) (*User, error)
	ServiceUpdateUsers(id string, input UpdateUsersRequest) (*User, error)
	ServiceDeleteUsers(id string) error
	ServiceProfileUsers(id float64) (*User, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) ServiceGetUsers() ([]User, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		log.Printf("Failed to get all users from repository: %v", err)
		return nil, appErrors.ErrInternal
	}
	return users, err
}

func (s *service) ServiceGetUsersByID(id string) (*User, error) {
	user, err := s.Repo.GetUsersByID(id)

	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			log.Printf("User with ID %s not found in database", id)
			return nil, appErrors.ErrNotFound
		}

		log.Printf("Failed to query user with ID %s: %v", id, err)
		return nil, appErrors.ErrInternal
	}

	return user, nil
}

func (s *service) ServiceCreateUsers(input CreateUsersRequest) (*User, error) {

	password, err := utils.HashedPassword(input.Password)
	if err != nil {
		log.Printf("Failed to hash password for user %s: %v", input.Email, err)
		return nil, appErrors.ErrInternal
	}

	user := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
		RoleID:   input.RoleID,
	}

	err = s.Repo.CreateUsers(&user)
	if err != nil {
		log.Printf("Failed to create user %s in database: %v", user.Email, err)
		return nil, appErrors.ErrInternal
	}

	return &user, err
}

func (s *service) ServiceUpdateUsers(id string, input UpdateUsersRequest) (*User, error) {
	user, err := s.Repo.GetUsersByID(id)
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			log.Printf("User with ID %s not found in database", id)
			return nil, appErrors.ErrNotFound
		}

		log.Printf("Failed to query user with ID %s: %v", id, err)
		return nil, appErrors.ErrInternal
	}

	password, err := utils.HashedPassword(input.Password)
	if err != nil {
		log.Printf("Failed to hash new password for user ID %s: %v", id, err)
		return nil, appErrors.ErrInternal
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = password
	user.RoleID = input.RoleID

	err = s.Repo.SaveUsers(user)
	if err != nil {
		log.Printf("Failed to save updated user with ID %s: %v", id, err)
		return nil, appErrors.ErrInternal
	}
	return user, nil
}

func (s *service) ServiceDeleteUsers(id string) error {
	user, err := s.Repo.GetUsersByID(id)
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			log.Printf("User with ID %s not found in database", id)
			return appErrors.ErrNotFound
		}

		log.Printf("Failed to query user with ID %s: %v", id, err)
		return appErrors.ErrInternal
	}

	err = s.Repo.DeleteUsers(user)
	if err != nil {
		log.Printf("Failed to delete user with ID %s from database: %v", id, err)
		return appErrors.ErrInternal
	}

	return nil
}

func (s *service) ServiceProfileUsers(id float64) (*User, error) {
	str_id := fmt.Sprintf("%.0f", id)
	user, err := s.Repo.GetUsersByID(str_id)
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			log.Printf("User with ID %s not found in database", str_id)
			return nil, appErrors.ErrNotFound
		}

		log.Printf("Failed to query user with ID %s: %v", str_id, err)
		return nil, appErrors.ErrInternal
	}
	return user, nil
}
