package models

import "gorm.io/gorm"

type AuditLog struct {
	gorm.Model
	CompanyID  uint
	UserID     uint
	Action     string
	Resource   string
	ResourceID string
	IP         string
}
