package testhelpers

import (
	"api-backend-saas/internal/models"
	"api-backend-saas/internal/security"

	"gorm.io/gorm"
)

func CreateTestUser(db *gorm.DB, email, password string) models.User {
	hashed, _ := security.HashPassword(password)

	user := models.User{
		Name:      "Teste",
		Email:     email,
		Password:  hashed,
		CompanyID: 1,
		RoleID:    1,
	}

	db.Create(&user)
	return user
}
