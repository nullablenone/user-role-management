package routes

import (
	"manajemen-user/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/user", controllers.Test)
	router.POST("/register", controllers.Register(db))

	return router
}
