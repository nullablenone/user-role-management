package user

import (
	"golang.org/x/crypto/bcrypt"
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
		return nil, err
	}
	return users, err
}

func (s *service) ServiceGetUsersByID(id string) (*User, error) {
	user, err := s.Repo.GetUsersByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) ServiceCreateUsers(input CreateUsersRequest) (*User, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(password),
		RoleID:   input.RoleID,
	}

	err = s.Repo.CreateUsers(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (s *service) ServiceUpdateUsers(id string, input UpdateUsersRequest) (*User, error) {
	user, err := s.Repo.GetUsersByID(id)
	if err != nil {
		return nil, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = string(password)
	user.RoleID = input.RoleID

	err = s.Repo.SaveUsers(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) ServiceDeleteUsers(id string) error {
	user, err := s.Repo.GetUsersByID(id)
	if err != nil {
		return err
	}

	err = s.Repo.DeleteUsers(user)
	if err != nil {
		return err
	}

	return nil
}
