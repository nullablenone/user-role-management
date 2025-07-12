package user

import (
	"manajemen-user/internal/domain/role"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	RoleID    uint
	Role      role.Role // Relasi ke model domain Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
