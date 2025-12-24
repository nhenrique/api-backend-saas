package database

import (
	"log"

	"github.com/nhenrique/api-backend-saas/internal/infra/persistence/gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=saas port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
	}
	DB = db

	// AutoMigrate: Cria as tabelas automaticamente com base nos modelos
	err = DB.AutoMigrate(
		&models.Company{},
		&models.Permission{},
		&models.Role{},
		&models.User{},
		&models.AuditLog{},
	)

	// Chama o Seed para adicionar dados iniciais
	Seed(db)

	if err != nil {
		log.Fatal("AutoMigrate failed")
	}
	log.Println("Database migration and seeding completed")
}
