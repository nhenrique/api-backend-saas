package user

import (
	"strings"

	"github.com/google/uuid"
)

func NewUser(
	name string,
	email string,
	hashedPassword string,
	companyID uuid.UUID,
	roleID uuid.UUID,
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

	if companyID == uuid.Nil {
		return nil, ErrInvalidCompany
	}

	if roleID == uuid.Nil {
		return nil, ErrInvalidRole
	}

	return &User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CompanyID: companyID,
		RoleID:    roleID,
	}, nil
}
