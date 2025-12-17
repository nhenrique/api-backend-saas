package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequirePermission(required string) gin.HandlerFunc {
	return func(c *gin.Context) {

		permsRaw, exists := c.Get("permissions")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "permissões não encontradas",
			})
			return
		}

		perms, ok := permsRaw.([]string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "formato de permissão inválido",
			})
			return
		}

		for _, p := range perms {
			if p == required {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "permissão negada",
		})
	}
}
