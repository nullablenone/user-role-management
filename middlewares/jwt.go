package middlewares

import (
	"manajemen-user/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token tidak boleh kosong",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Format token tidak valid, harus diawali dengan 'Bearer '",
			})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

		c.Set("claims", claims)
		c.Next()
	}
}
