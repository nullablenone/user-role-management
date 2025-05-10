package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireRole(expectedRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mapClaims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "gagal mengambil claims",
			})
			c.Abort()
			return
		}

		claims, ok := mapClaims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "gagal invalid token claims",
			})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "tidak menemukan role",
			})
			c.Abort()
			return
		}

		if role != expectedRole {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "akses di tolak",
			})
			c.Abort()
			return
		}

	}
}
