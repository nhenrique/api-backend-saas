package gormrepo

import (
	"context"

	domain "github.com/nhenrique/api-backend-saas/internal/domain/user"
	"github.com/nhenrique/api-backend-saas/internal/infra/persistence/gorm/models"
)

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	model := models.UserModel{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CompanyID: user.CompanyID,
		RoleID:    user.RoleID,
	}

	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var model models.UserModel

	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&model).Error; err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		CompanyID: model.CompanyID,
		RoleID:    model.RoleID,
	}, nil
}
