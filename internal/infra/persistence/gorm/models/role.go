package models

type Role struct {
	BaseModel
	Name        string       `gorm:"uniqueIndex"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
