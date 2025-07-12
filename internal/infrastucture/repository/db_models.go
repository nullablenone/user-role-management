package repository

import "gorm.io/gorm"

// RoleModel adalah representasi tabel 'roles' untuk GORM.
type RoleModel struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Deskription string `gorm:"type:varchar(50);not null"`
}

func (RoleModel) TableName() string {
	return "roles"
}

// UserModel adalah representasi tabel 'users' untuk GORM.
type UserModel struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(50);not null"`
	Email    string    `gorm:"type:varchar(350);uniqueIndex;not null"`
	Password string    `gorm:"type:varchar(350);not null"`
	RoleID   uint      `gorm:"default:1"`
	Role     RoleModel `gorm:"foreignKey:RoleID"`
}

func (UserModel) TableName() string {
	return "users"
}
