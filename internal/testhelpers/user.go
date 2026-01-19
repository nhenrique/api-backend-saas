package testhelpers

import (
	"github.com/google/uuid"
	"github.com/nhenrique/api-backend-saas/internal/infra/persistence/gorm/models"
	"github.com/nhenrique/api-backend-saas/internal/security"
	"gorm.io/gorm"
)

func CreateTestUser(db *gorm.DB, email, password string) models.User {
	hashed, _ := security.HashPassword(password)

	user := models.User{
		BaseModel: models.BaseModel{
			ID: uuid.New(),
		},
		Name:      "Teste",
		Email:     email,
		Password:  hashed,
		CompanyID: uuid.New(),
		RoleID:    uuid.New(),
	}

	db.Create(&user)
	return user
}
