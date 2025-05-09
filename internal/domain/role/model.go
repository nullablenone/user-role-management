package role

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Deskription string `gorm:"type:varchar(50);not null"`
}
