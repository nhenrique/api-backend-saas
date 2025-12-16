package database

import (
	"log"

	"api-backend-saas/internal/models"
	"api-backend-saas/internal/security"

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
		&models.Plan{},
		&models.User{},
		&models.AuditLog{},
	)

	// Chama o Seed para adicionar dados iniciais
	Seed()

	if err != nil {
		log.Fatal("AutoMigrate failed")
	}
	log.Println("Database migration and seeding completed")
}

func Seed() {
	// Verificar se já existe algum admin, senão cria um novo.
	var adminUser models.User
	if err := DB.Where("email = ?", "admin@admin.com").First(&adminUser).Error; err != nil {
		// Criar plano padrão
		plan := models.Plan{
			Name:      "Basic",
			RateLimit: 100,
		}
		DB.Create(&plan)

		// Criar empresa de exemplo
		company := models.Company{
			Name:   plan.Name,
			PlanID: plan.ID,
		}
		DB.Create(&company)

		hashedPassword, _ := security.HashPassword("123456")

		// Criar usuário admin
		adminUser = models.User{
			Name:      "Admin User",
			Email:     "admin@admin.com",
			Password:  hashedPassword,
			CompanyID: company.ID,
			Role:      "admin",
		}
		DB.Create(&adminUser)

		log.Println("Seed data inserted successfully")
	} else {
		log.Println("Admin user already exists")
	}
}
