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
	admin.GET("/users/:id", controllers.GetUsersByID(db))
	admin.POST("/users", controllers.CreateUsers(db))
	admin.PUT("/users/:id", controllers.UpdateUsers(db))
	admin.DELETE("/users/:id", controllers.DeleteUsers(db))

	return router
}
