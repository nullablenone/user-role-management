package main

import (
	"log"
	"manajemen-user/config"
	"manajemen-user/internal/domain/auth"
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
	authService := auth.NewService(userRepo)
	userService := user.NewService(userRepo)
	roleService := role.NewService(roleRepo)
	// Set User Hundler
	authHandler := auth.NewHandler(authService)
	userHandler := user.NewHandler(userService)
	roleHandler := role.NewHandler(roleService)

	router := routes.SetupRoutes(authHandler, userHandler, roleHandler)
	router.Run(":8080")
}
