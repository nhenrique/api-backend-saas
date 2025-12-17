package middlewares

import (
	"net/http"
	"strings"

	"api-backend-saas/internal/config"
	"api-backend-saas/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// =========================
		// Header Authorization
		// =========================
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

		// =========================
		// Parse JWT
		// =========================
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

		// =========================
		// Extrai dados do JWT
		// =========================
		userID := uint(claims["sub"].(float64))
		role := claims["role"].(string)
		companyID := uint(claims["company_id"].(float64))

		// =========================
		// Carrega permissões do banco
		// =========================
		var permissions []string

		err = database.DB.
			Table("permissions").
			Select("permissions.name").
			Joins("JOIN role_permissions rp ON rp.permission_id = permissions.id").
			Joins("JOIN roles r ON r.id = rp.role_id").
			Joins("JOIN users u ON u.role_id = r.id").
			Where("u.id = ?", userID).
			Scan(&permissions).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "erro ao carregar permissões",
			})
			return
		}

		// =========================
		// Injeta no contexto
		// =========================
		c.Set("user_id", userID)
		c.Set("role", role)
		c.Set("company_id", companyID)
		c.Set("permissions", permissions)

		c.Next()
	}
}
