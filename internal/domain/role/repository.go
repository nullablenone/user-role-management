package role

type Repository interface {
	GetAllRoles() ([]Role, error)
	GetRolesByID(id string) (*Role, error)
	CreateRoles(role *Role) error
	SaveRoles(role *Role) error
	DeleteRoles(role *Role) error
}
