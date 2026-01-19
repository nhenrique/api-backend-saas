package testhelpers

import (
	"context"

	"github.com/nhenrique/api-backend-saas/internal/domain/user"
)

type UserRepositoryMock struct {
	CreateFn        func(ctx context.Context, user *user.User) error
	FindByCompanyFn func(ctx context.Context, companyID uint) ([]user.User, error)
}

func (m *UserRepositoryMock) Create(ctx context.Context, u *user.User) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, u)
	}
	return nil
}

func (m *UserRepositoryMock) FindByCompany(ctx context.Context, companyID uint) ([]user.User, error) {
	if m.FindByCompanyFn != nil {
		return m.FindByCompanyFn(ctx, companyID)
	}
	return []user.User{}, nil
}
