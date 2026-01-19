package converters

import (
	"github.com/google/uuid"

	domain "github.com/nhenrique/api-backend-saas/internal/domain/user"
	"github.com/nhenrique/api-backend-saas/internal/infra/persistence/gorm/models"
)

func UserToModel(u *domain.User) (*models.User, error) {
	var id uuid.UUID
	var err error

	if u.ID != "" {
		id, err = uuid.Parse(u.ID)
		if err != nil {
			return nil, err
		}
	}

	return &models.User{
		BaseModel: models.BaseModel{
			ID: id,
		},
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CompanyID: uuid.MustParse(u.CompanyID),
		RoleID:    uuid.MustParse(u.RoleID),
	}, nil
}

func ModelToUser(m *models.User) *domain.User {
	return &domain.User{
		ID:        m.ID.String(),
		Name:      m.Name,
		Email:     m.Email,
		Password:  m.Password,
		CompanyID: m.CompanyID.String(),
		RoleID:    m.RoleID.String(),
	}
}
