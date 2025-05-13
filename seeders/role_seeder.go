package seeders

import (
	"manajemen-user/internal/domain/role"

	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []role.Role{
		{Name: "user", Deskription: "Limited access to personal features only"},
		{Name: "admin", Deskription: "Full access to all system features"},
	}

	for _, r := range roles {
		db.FirstOrCreate(&r, role.Role{Name: r.Name})
	}
}
