package middlewares

import (
	"net/http"
	"strings"

	"github.com/nhenrique/api-backend-saas/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token não informado",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "formato de token inválido",
			})
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return config.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido ou expirado",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "claims inválidas",
			})
			return
		}

		// ✅ SAFE ASSERTIONS
		userID, _ := claims["sub"].(float64)
		companyID, _ := claims["company_id"].(float64)

		var permissions []string
		if p, ok := claims["permissions"].([]interface{}); ok {
			for _, v := range p {
				if s, ok := v.(string); ok {
					permissions = append(permissions, s)
				}
			}
		}

		// 🔑 Inject context
		c.Set("user_id", uint(userID))
		c.Set("role", claims["role"])
		c.Set("company_id", uint(companyID))
		c.Set("permissions", permissions)

		c.Next()
	}
}
