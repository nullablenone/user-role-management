package utils

import (
	"errors"
	appErrors "manajemen-user/internal/errors"
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

func HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, appErrors.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Resource not found"})
	case errors.Is(err, appErrors.ErrInternal):
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An internal server error occurred"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An unknown error occurred"})
	}
}
