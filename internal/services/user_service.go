package services

import (
	"context"

	userdomain "github.com/nhenrique/api-backend-saas/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo userdomain.Repository
}

func NewUserService(userRepo userdomain.Repository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	email string,
	password string,
	companyID string,
	roleID string,
) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := userdomain.NewUser(
		name,
		email,
		string(hashed),
		companyID,
		roleID,
	)
	if err != nil {
		return err
	}

	return s.userRepo.Create(ctx, user)
}
