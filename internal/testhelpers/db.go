package testhelpers

import (
	"api-backend-saas/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	db.AutoMigrate(
		&models.Company{},
		&models.Role{},
		&models.Permission{},
		&models.User{},
	)

	return db
}
