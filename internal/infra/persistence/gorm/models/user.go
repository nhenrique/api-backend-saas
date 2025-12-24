package models

import (
	"github.com/google/uuid"
)

type User struct {
	BaseModel
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CompanyID uuid.UUID
	Company   Company
	RoleID    uuid.UUID
	Role      Role
}
