package routes

import (
	"manajemen-user/controllers"
	"manajemen-user/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/user", controllers.Test)
	router.POST("/register", controllers.Register(db))
	router.POST("/login", controllers.Login(db))

	admin := router.Group("/admin", middlewares.JWTMiddleware())
	admin.GET("/users", controllers.GetUsers(db))
	// admin.POST("/users/create", controllers.CreateUsers(db))

	return router
}
