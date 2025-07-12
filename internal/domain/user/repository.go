package user

type Repository interface {
	GetAllUsers() ([]User, error)
	GetUsersByID(id string) (*User, error)
	CreateUsers(user *User) error
	SaveUsers(user *User) error
	DeleteUsers(user *User) error
	FindByEmailWithRole(email string) (*User, error)
}
