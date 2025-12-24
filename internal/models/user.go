package models

import "github.com/google/uuid"

type UserModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CompanyID uuid.UUID `gorm:"type:uuid;index"`
	RoleID    uuid.UUID `gorm:"type:uuid"`
	Role      Role
}
