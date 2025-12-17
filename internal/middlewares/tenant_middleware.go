package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnforceTenant() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenant := c.GetUint("company_id")

		if tenant == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "tenant inválido",
			})
			return
		}

		c.Set("tenant_id", tenant)
		c.Next()
	}
}
