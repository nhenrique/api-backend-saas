package models

import "time"

type AuditLog struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	CompanyID uint
	Action    string
	Path      string
	Method    string
	CreatedAt time.Time
}
