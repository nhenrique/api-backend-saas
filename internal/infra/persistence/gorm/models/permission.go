package models

type Permission struct {
	BaseModel
	Name string `gorm:"uniqueIndex"`
}
