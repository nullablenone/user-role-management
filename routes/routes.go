package routes

import (
	"manajemen-user/controllers"
	"manajemen-user/internal/domain/user"
	"manajemen-user/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, userHandler *user.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/user", controllers.Test)
	router.POST("/register", controllers.Register(db))
	router.POST("/login", controllers.Login(db))

	admin := router.Group("/admin", middlewares.JWTMiddleware())
	admin.GET("/users", userHandler.GetUsers)
	admin.GET("/users/:id", userHandler.GetUsersByID)
	admin.POST("/users", userHandler.CreateUsers)
	admin.PUT("/users/:id", userHandler.UpdateUsers)
	admin.DELETE("/users/:id", userHandler.DeleteUsers)

	return router
}
