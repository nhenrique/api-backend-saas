package user

import (
	"strings"

	"github.com/nhenrique/api-backend-saas/internal/domain/id/id"
)

func NewUser(
	name string,
	email string,
	hashedPassword string,
	companyID string,
	roleID string,
) (*User, error) {

	if strings.TrimSpace(name) == "" {
		return nil, ErrInvalidName
	}

	if strings.TrimSpace(email) == "" {
		return nil, ErrInvalidEmail
	}

	if hashedPassword == "" {
		return nil, ErrInvalidPassword
	}

	if companyID == "" {
		return nil, ErrInvalidCompany
	}

	if roleID == "" {
		return nil, ErrInvalidRole
	}

	return &User{
		ID:        id.New(),
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CompanyID: companyID,
		RoleID:    roleID,
	}, nil
}
