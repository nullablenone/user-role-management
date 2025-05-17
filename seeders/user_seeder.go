package seeders

import (
	"log"
	"manajemen-user/internal/domain/user"
	"manajemen-user/utils"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	users := []user.User{
		{Name: "admin", Email: "admin@gmail.com", Password: "admin@gmail.com", RoleID: 2},
		{Name: "yesaGG", Email: "yesaGG@gmail.com", Password: "yesaGG@gmail.com", RoleID: 1},
		{Name: "adam", Email: "adam@gmail.com", Password: "adam@gmail.com", RoleID: 1},
		{Name: "ahmad", Email: "ahmad@gmail.com", Password: "ahmad@gmail.com", RoleID: 1},
	}

	for _, u := range users {
		password, err := utils.HashedPassword(u.Password)
		if err != nil {
			log.Printf("Failed hashing password for user %s: %v", u.Email, err)
			continue
		}

		u.Password = password
		db.FirstOrCreate(&u, user.User{Email: u.Email})
	}
}
