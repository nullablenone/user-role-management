package repository

import (
	"time"
)

// RoleModel adalah representasi tabel 'roles' untuk GORM.
type RoleModel struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Deskription string `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (RoleModel) TableName() string {
	return "roles"
}

// UserModel adalah representasi tabel 'users' untuk GORM.
type UserModel struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"type:varchar(50);not null"`
	Email     string    `gorm:"type:varchar(350);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(350);not null"`
	RoleID    uint      `gorm:"default:1"`
	Role      RoleModel `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}
