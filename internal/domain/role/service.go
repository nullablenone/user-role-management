package role

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
	var role []Role
	roles, err := s.Repo.GetAllRoles(role)
	if err != nil {
		return nil, err
	}
	return roles, err
}

func (s *service) ServiceGetRolesByID(id string) (*Role, error) {
	role, err := s.Repo.GetRolesByID(id)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &role, err
}

func (s *service) ServiceUpdateRoles(id string, input UpdateRolesRequest) (*Role, error) {
	role, err := s.Repo.GetRolesByID(id)
	if err != nil {
		return nil, err
	}

	role.Name = input.Name
	role.Deskription = input.Deskription

	err = s.Repo.SaveRoles(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *service) ServiceDeleteRoles(id string) error {
	role, err := s.Repo.GetRolesByID(id)
	if err != nil {
		return err
	}

	err = s.Repo.DeleteRoles(role)
	if err != nil {
		return err
	}

	return nil
}
