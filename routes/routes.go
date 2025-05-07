package routes

import (
	"manajemen-user/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/user", controllers.Test)

	return router
}
