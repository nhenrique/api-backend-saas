package models

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	Name      string
	RateLimit int
}
