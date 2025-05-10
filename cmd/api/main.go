package main

import (
	"log"
	"manajemen-user/config"
	"manajemen-user/internal/domain/role"
	"manajemen-user/internal/domain/user"
	"manajemen-user/routes"
)

func main() {
	// Env Load
	config.LoadENV()
	// Connect DB
	db := config.ConnectDB()
	// Migrate Table
	err := db.AutoMigrate(role.Role{}, user.User{})
	if err != nil {
		log.Fatal(err)
	}
	// Set Repo
	userRepo := user.NewRepository(db)
	roleRepo := role.NewRepository(db)
	// Set service
	userService := user.NewService(userRepo)
	roleService := role.NewService(roleRepo)
	// Set User Hundler
	userHandler := user.NewHandler(userService)
	roleHandler := role.NewHandler(roleService)

	router := routes.SetupRoutes(db, userHandler, roleHandler)
	router.Run(":8080")
}
