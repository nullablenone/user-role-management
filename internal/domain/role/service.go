package role

import (
	"errors"
	"log"
	appErrors "manajemen-user/internal/errors"
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
		log.Printf("Failed to get all roles from repository: %v", err)
		return nil, appErrors.ErrInternal
	}

	return roles, err
}

func (s *service) ServiceGetRolesByID(id string) (*Role, error) {
	role, err := s.Repo.GetRolesByID(id)

	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			log.Printf("Role with ID %s not found in database", id)
			return nil, appErrors.ErrNotFound
		}

		log.Printf("Failed to query role with ID %s: %v", id, err)
		return nil, appErrors.ErrInternal
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
		log.Printf("Failed to create role '%s' in database: %v", role.Name, err)
		return nil, appErrors.ErrInternal
	}

	return &role, err
}

func (s *service) ServiceUpdateRoles(id string, input UpdateRolesRequest) (*Role, error) {
	role, err := s.Repo.GetRolesByID(id)

	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			log.Printf("Role with ID %s not found in database", id)
			return nil, appErrors.ErrNotFound
		}

		log.Printf("Failed to query role with ID %s: %v", id, err)
		return nil, appErrors.ErrInternal
	}

	role.Name = input.Name
	role.Deskription = input.Deskription

	err = s.Repo.SaveRoles(role)
	if err != nil {
		log.Printf("Failed to save updated role with ID %s: %v", id, err)
		return nil, appErrors.ErrInternal
	}
	return role, nil
}

func (s *service) ServiceDeleteRoles(id string) error {
	role, err := s.Repo.GetRolesByID(id)
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			log.Printf("Role with ID %s not found in database", id)
			return appErrors.ErrNotFound
		}

		log.Printf("Failed to query role with ID %s: %v", id, err)
		return appErrors.ErrInternal
	}

	err = s.Repo.DeleteRoles(role)
	if err != nil {
		log.Printf("Failed to delete role with ID %s from database: %v", id, err)
		return appErrors.ErrInternal
	}

	return nil
}
