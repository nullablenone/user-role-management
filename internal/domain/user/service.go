package user

import (
	"errors"
	"fmt"
	"manajemen-user/utils"

	"gorm.io/gorm"
)

type Service interface {
	ServiceGetUsers() ([]User, error)
	ServiceGetUsersByID(id string) (*User, error)
	ServiceCreateUsers(input CreateUsersRequest) (*User, error)
	ServiceUpdateUsers(id string, input UpdateUsersRequest) (*User, error)
	ServiceDeleteUsers(id string) error
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) ServiceGetUsers() ([]User, error) {
	var user []User
	users, err := s.Repo.GetAllUsers(user)
	if err != nil {
		return nil, fmt.Errorf("ServiceGetUsers: failed to get users: %w", err)
	}
	return users, err
}

func (s *service) ServiceGetUsersByID(id string) (*User, error) {
	user, err := s.Repo.GetUsersByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ServiceGetUsersByID: user with ID %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("ServiceGetUsersByID: failed to get user: %w", err)
	}
	return user, nil
}

func (s *service) ServiceCreateUsers(input CreateUsersRequest) (*User, error) {

	password, err := utils.HashedPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("ServiceCreateUsers: failed hashed password: %w", err)
	}

	user := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
		RoleID:   input.RoleID,
	}

	err = s.Repo.CreateUsers(&user)
	if err != nil {
		return nil, fmt.Errorf("ServiceCreateUsers: failed to create user: %w", err)
	}

	return &user, err
}

func (s *service) ServiceUpdateUsers(id string, input UpdateUsersRequest) (*User, error) {
	user, err := s.Repo.GetUsersByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ServiceUpdateUsers: user with ID %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("ServiceUpdateUsers: failed to get user: %w", err)
	}

	password, err := utils.HashedPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = password
	user.RoleID = input.RoleID

	err = s.Repo.SaveUsers(user)
	if err != nil {
		return nil, fmt.Errorf("ServiceUpdateUsers: failed to update user: %w", err)
	}
	return user, nil
}

func (s *service) ServiceDeleteUsers(id string) error {
	user, err := s.Repo.GetUsersByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ServiceDeleteUsers: User with ID %s not found: %w", id, err)
		}
		return fmt.Errorf("ServiceDeleteUsers: failed to get user: %w", err)
	}

	err = s.Repo.DeleteUsers(user)
	if err != nil {
		return fmt.Errorf("ServiceDeleteUsers: failed to delete user: %w", err)
	}

	return nil
}
