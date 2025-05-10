package middlewares

import (
	"fmt"
	"manajemen-user/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(expectedRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mapClaims, exists := c.Get("claims")
		if !exists {
			utils.RespondError(c, http.StatusUnauthorized, "RequireRole: failed to take claims")
			c.Abort()
			return
		}

		claims, err := utils.AssertTypeClaims(mapClaims)
		if err != nil {
			utils.RespondError(c, http.StatusUnauthorized, fmt.Sprintf("RequireRole: %v", err))
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			utils.RespondError(c, http.StatusUnauthorized, "RequireRole: did not find role")
			c.Abort()
			return
		}

		if role != expectedRole {
			utils.RespondError(c, http.StatusUnauthorized, "RequireRole: access denied")
			c.Abort()
			return
		}

	}
}
