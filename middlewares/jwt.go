package middlewares

import (
	"fmt"
	"manajemen-user/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.RespondError(c, http.StatusUnauthorized, "token cannot be empty")
			c.Abort()
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			utils.RespondError(c, http.StatusUnauthorized, "invalid token format, must be started with 'Bearer '")

			c.Abort()
			return
		}

		token := strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := utils.ValidateToken(token)
		if err != nil {
			utils.RespondError(c, http.StatusUnauthorized, fmt.Sprintf("JWTMiddleware: %v", err))
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
