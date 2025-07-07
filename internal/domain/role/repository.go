package role

type Repository interface {
	GetAllRoles(user []Role) ([]Role, error)
	GetRolesByID(id string) (*Role, error)
	CreateRoles(user *Role) error
	SaveRoles(user *Role) error
	DeleteRoles(user *Role) error
}
