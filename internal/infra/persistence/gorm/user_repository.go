package gorm

import (
	"github.com/google/uuid"
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
	userID, err := uuid.Parse(user.ID)
	if err != nil {
		return err
	}

	companyID, err := uuid.Parse(user.CompanyID)
	if err != nil {
		return err
	}

	roleID, err := uuid.Parse(user.RoleID)
	if err != nil {
		return err
	}

	model := models.User{
		BaseModel: models.BaseModel{
			ID: userID,
		},
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CompanyID: companyID,
		RoleID:    roleID,
	}

	return r.db.Create(&model).Error
}
