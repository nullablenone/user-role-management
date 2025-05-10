package user

import (
	"manajemen-user/internal/domain/role"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);not null"`
	Email    string `gorm:"type:varchar(350);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(350);not null"`
	RoleID   uint   `gorm:"default:1"` // Foreign key ke Role
	Role     role.Role
}
