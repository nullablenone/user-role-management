package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  "error",
		"message": message,
	})
}

func RespondSuccess(c *gin.Context, data any, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}
