package main

import (
	"log"
	"manajemen-user/config"
	"manajemen-user/internal/domain/auth"
	"manajemen-user/internal/domain/role"
	"manajemen-user/internal/domain/user"
	"manajemen-user/routes"
	"manajemen-user/seeders"

	docs "manajemen-user/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @description  This API allows managing users and roles with authentication.
// @description
// @description  🔐 **Authorization**
// @description  To access protected endpoints, click "Authorize" and enter your token in this format:
// @description  `Bearer <your-token>` (with a space after Bearer).
// @description
// @description  👤 **Login as Admin**
// @description  If you want to log in as an admin, please contact the developer. And that's me.  
// @description  **Contact Email**: useryesa9@gmail.com
// @description
// @description  Make sure to copy the token from the login response and prepend it with `Bearer ` before pasting it into the Authorize box.

// @contact.name   Developer Support
// @contact.email  useryesa9@gmail.com

// @tag.name Auth
// @tag.description Endpoint to login and generate token

// @tag.name User - Profile
// @tag.description Endpoint for regular users to access their profile

// @tag.name Admin - Users
// @tag.description Endpoint for admin in managing user data

// @tag.name Admin - Roles
// @tag.description Endpoint for admin to manage roles

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	//swag config
	docs.SwaggerInfo.Title = "User Role Management API"

	// Env Load
	config.LoadENV()

	// Connect DB

	db := config.ConnectDB()
	// Migrate Table

	err := db.AutoMigrate(role.Role{}, user.User{})
	if err != nil {
		log.Fatal(err)
	}

	// Call Seeder
	seeders.SeedRole(db)
	seeders.SeedUser(db)

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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
