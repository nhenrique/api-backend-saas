package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CompanyID uuid.UUID
	RoleID    uuid.UUID
}
