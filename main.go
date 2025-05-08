package main

import (
	"log"
	"manajemen-user/config"
	"manajemen-user/models"
	"manajemen-user/routes"
)

func main() {
	config.LoadENV()
	db := config.ConnectDB()
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
		return
	}

	router := routes.SetupRoutes(db)
	router.Run(":8080")
}
