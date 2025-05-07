package main

import (
	"manajemen-user/config"
	"manajemen-user/routes"
)

func main() {
	config.LoadENV()
	config.ConnectDB()

	router := routes.SetupRoutes()
	router.Run(":8080")
}
