package routes

import (
	"manajemen-user/controllers"
	"manajemen-user/internal/domain/role"
	"manajemen-user/internal/domain/user"
	"manajemen-user/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, userHandler *user.Handler, roleHandler *role.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/user", controllers.Test)
	router.POST("/register", controllers.Register(db))
	router.POST("/login", controllers.Login(db))

	admin := router.Group("/admin", middlewares.JWTMiddleware(), middlewares.RequireRole("admin"))
	admin.GET("/users", userHandler.GetUsers)
	admin.GET("/users/:id", userHandler.GetUsersByID)
	admin.POST("/users", userHandler.CreateUsers)
	admin.PUT("/users/:id", userHandler.UpdateUsers)
	admin.DELETE("/users/:id", userHandler.DeleteUsers)

	admin.GET("/roles", roleHandler.GetRoles)
	admin.GET("/roles/:id", roleHandler.GetRolesByID)
	admin.POST("/roles", roleHandler.CreateRoles)
	admin.PUT("/roles/:id", roleHandler.UpdateRoles)
	admin.DELETE("/roles/:id", roleHandler.DeleteRoles)

	return router
}
