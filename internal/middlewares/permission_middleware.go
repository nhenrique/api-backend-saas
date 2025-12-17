package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequirePermission(required string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// =========================
		// Busca permissões no contexto
		// =========================
		value, exists := c.Get("permissions")
		if !exists || value == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "permissões não carregadas no contexto",
			})
			return
		}

		perms, ok := value.([]string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "formato inválido de permissões",
			})
			return
		}

		// =========================
		// Valida permissão exigida
		// =========================
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
