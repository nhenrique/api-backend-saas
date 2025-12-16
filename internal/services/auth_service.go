package services

import (
	"errors"
	"time"

	"api-backend-saas/internal/config"
	"api-backend-saas/internal/models"
	"api-backend-saas/internal/security"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User

	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("credenciais inválidas")
	}

	if !security.CheckPassword(password, user.Password) {
		return "", errors.New("credenciais inválidas")
	}

	claims := jwt.MapClaims{
		"sub":        user.ID,
		"email":      user.Email,
		"role":       user.Role,
		"company_id": user.CompanyID,
		"iss":        config.JWTIssuer,
		"exp":        time.Now().Add(config.JWTExpireDuration()).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(config.JWTSecret)
	if err != nil {
		return "", err
	}

	return signed, nil
}
