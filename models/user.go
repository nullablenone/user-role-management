package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);not null"`
	Email    string `gorm:"type:varchar(350);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(350);not null"`
	Role     string `gorm:"type:varchar(50);default:'user';not null"`
}
