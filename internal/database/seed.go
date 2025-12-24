package database

import (
	"log"

	"github.com/nhenrique/api-backend-saas/internal/models"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	log.Println("🌱 Running database seeds...")

	seedCompanies(db)
	seedPermissions(db)
	seedRoles(db)
	seedRolesAndPermissions(db)
	seedAdminUser(db)

	log.Println("✅ Seed completed")
}

func seedCompanies(db *gorm.DB) {
	company := models.Company{Name: "Default Company"}
	db.FirstOrCreate(&company, models.Company{Name: company.Name})
}

func seedPermissions(db *gorm.DB) {
	perms := []models.Permission{
		{Name: "user:create"},
		{Name: "user:view"},
		{Name: "user:delete"},
		{Name: "report:view"},
	}

	for _, p := range perms {
		db.FirstOrCreate(&p, models.Permission{Name: p.Name})
	}
}

func seedRoles(db *gorm.DB) {
	roles := []string{"admin", "gerente", "usuario"}

	for _, r := range roles {
		db.FirstOrCreate(&models.Role{}, models.Role{Name: r})
	}
}

func seedRolesAndPermissions(db *gorm.DB) {

	// =====================
	// Permissions
	// =====================
	perms := []models.Permission{
		{Name: "user:create"},
		{Name: "user:list"},
		{Name: "user:update"},
	}

	for _, p := range perms {
		db.FirstOrCreate(&p, models.Permission{Name: p.Name})
	}

	// =====================
	// Roles
	// =====================
	var adminRole models.Role
	db.FirstOrCreate(&adminRole, models.Role{Name: "admin"})

	var userCreate models.Permission
	db.Where("name = ?", "user:create").First(&userCreate)

	db.Model(&adminRole).Association("Permissions").Append(&userCreate)
}

func seedAdminUser(db *gorm.DB) {
	var adminRole models.Role
	db.Where("name = ?", "admin").First(&adminRole)

	var company models.Company
	db.First(&company)

	var admin models.User
	if err := db.Where("email = ?", "admin@admin.com").First(&admin).Error; err == nil {
		log.Println("👤 Admin already exists")
		return
	}

	admin = models.User{
		Name:      "Admin",
		Email:     "admin@admin.com",
		Password:  "1234",
		CompanyID: company.ID,
		RoleID:    adminRole.ID,
	}

	db.Create(&admin)
}
