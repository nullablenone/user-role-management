package main

import (
	"manajemen-user/config"
	"manajemen-user/internal/domain/user"
	"manajemen-user/routes"
)

func main() {
	// Env Load
	config.LoadENV()
	// Connect DB
	db := config.ConnectDB()
	// Migrate Table
	// if err := db.AutoMigrate(&models.User{}); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// Set Repo
	repo := user.NewRepository(db)
	// Set service
	service := user.NewService(repo)
	// Set User Hundler
	userHandler := user.NewHandler(service)

	router := routes.SetupRoutes(db, userHandler)
	router.Run(":8080")
}
