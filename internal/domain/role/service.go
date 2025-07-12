package role

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Service interface {
	ServiceGetRoles() ([]Role, error)
	ServiceGetRolesByID(id string) (*Role, error)
	ServiceCreateRoles(input CreateRolesRequest) (*Role, error)
	ServiceUpdateRoles(id string, input UpdateRolesRequest) (*Role, error)
	ServiceDeleteRoles(id string) error
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) ServiceGetRoles() ([]Role, error) {
	roles, err := s.Repo.GetAllRoles()
	if err != nil {
		return nil, fmt.Errorf("ServiceGetRoles: failed to get roles: %w", err)
	}
	return roles, err
}

func (s *service) ServiceGetRolesByID(id string) (*Role, error) {
	role, err := s.Repo.GetRolesByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ServiceGetRolesByID: role with ID %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("ServiceGetRolesByID: failed to get role: %w", err)
	}
	return role, nil
}

func (s *service) ServiceCreateRoles(input CreateRolesRequest) (*Role, error) {

	role := Role{
		Name:        input.Name,
		Deskription: input.Deskription,
	}

	err := s.Repo.CreateRoles(&role)
	if err != nil {
		return nil, fmt.Errorf("ServiceCreateRoles: failed to create role: %w", err)
	}

	return &role, err
}

func (s *service) ServiceUpdateRoles(id string, input UpdateRolesRequest) (*Role, error) {
	role, err := s.Repo.GetRolesByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ServiceUpdateRoles: role with ID %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("ServiceUpdateRoles: failed to get role: %w", err)
	}

	role.Name = input.Name
	role.Deskription = input.Deskription

	err = s.Repo.SaveRoles(role)
	if err != nil {
		return nil, fmt.Errorf("ServiceUpdateRoles: failed to update role: %w", err)
	}
	return role, nil
}

func (s *service) ServiceDeleteRoles(id string) error {
	role, err := s.Repo.GetRolesByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ServiceDeleteRoles: role with ID %s not found: %w", id, err)
		}
		return fmt.Errorf("ServiceDeleteRoles: failed to get role: %w", err)
	}

	err = s.Repo.DeleteRoles(role)
	if err != nil {
		return fmt.Errorf("ServiceDeleteRoles: failed to delete role: %w", err)
	}

	return nil
}
