package gorm

import (
	domain "github.com/nhenrique/api-backend-saas/internal/domain/user"
	"github.com/nhenrique/api-backend-saas/internal/infra/persistence/gorm/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	model := models.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CompanyID: user.CompanyID,
		RoleID:    user.RoleID,
	}
	return r.db.Create(&model).Error
}
