package services

import (
	"errors"
	"time"

	"github.com/nhenrique/api-backend-saas/internal/config"
	"github.com/nhenrique/api-backend-saas/internal/models"
	"github.com/nhenrique/api-backend-saas/internal/security"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User

	// 🔹 Carrega User + Role + Permissions
	if err := s.DB.
		Preload("Role.Permissions").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return "", errors.New("credenciais inválidas")
	}

	// 🔹 Verifica senha (hash)
	if !security.CheckPassword(password, user.Password) {
		return "", errors.New("credenciais inválidas")
	}

	// 🔹 Extrai permissions para string slice
	var permissions []string
	for _, p := range user.Role.Permissions {
		permissions = append(permissions, p.Name)
	}

	// 🔹 JWT Claims
	claims := jwt.MapClaims{
		"sub":         user.ID,
		"email":       user.Email,
		"role":        user.Role.Name, // ✅ string, não struct
		"permissions": permissions,    // ✅ RBAC real
		"company_id":  user.CompanyID,
		"iss":         config.JWTIssuer,
		"exp":         time.Now().Add(config.JWTExpireDuration()).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(config.JWTSecret)
	if err != nil {
		return "", err
	}

	return signed, nil
}
