package testhelpers

import (
	"time"

	"github.com/nhenrique/api-backend-saas/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type TestJWTClaims struct {
	UserID      uint
	Role        string
	CompanyID   uint
	Permissions []string
}

func GenerateTestJWT(data TestJWTClaims) string {
	claims := jwt.MapClaims{
		"sub":         data.UserID,
		"role":        data.Role,
		"company_id":  data.CompanyID,
		"permissions": data.Permissions,
		"iss":         "test",
		"exp":         time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString(config.JWTSecret)

	return signed
}
