package database

import (
	"log"

	"github.com/google/uuid"
	"github.com/nhenrique/api-backend-saas/internal/infra/persistence/gorm/models"
	"github.com/nhenrique/api-backend-saas/internal/security"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	log.Println("🌱 Running database seeds...")

	seedPlans(db)
	seedCompanies(db)
	seedPermissions(db)
	seedRoles(db)
	seedRolesAndPermissions(db)
	seedAdminUser(db)

	log.Println("✅ Seed completed")
}

func seedPlans(db *gorm.DB) models.Plan {
	var plan models.Plan
	if err := db.Where("name = ?", "default").First(&plan).Error; err == nil {
		return plan
	}

	plan = models.Plan{
		BaseModel: models.BaseModel{
			ID: uuid.New(),
		},
		Name:      "default",
		RateLimit: 0,
	}
	db.Create(&plan)
	return plan
}

func seedCompanies(db *gorm.DB) {
	var company models.Company

	err := db.Where("name = ?", "Default Company").First(&company).Error
	if err == nil {
		return
	}

	var plan models.Plan
	db.First(&plan)

	company = models.Company{
		BaseModel: models.BaseModel{
			ID: uuid.New(),
		},
		Name:   "Default Company",
		PlanID: plan.ID,
	}

	db.Create(&company)
}

func seedPermissions(db *gorm.DB) {
	perms := []string{
		"user:create",
		"user:view",
		"user:delete",
		"report:view",
	}

	for _, name := range perms {
		var perm models.Permission

		err := db.Where("name = ?", name).First(&perm).Error
		if err == nil {
			continue
		}

		db.Create(&models.Permission{
			BaseModel: models.BaseModel{
				ID: uuid.New(),
			},
			Name: name,
		})
	}
}

func seedRoles(db *gorm.DB) {
	roles := []string{"admin", "gerente", "usuario"}

	for _, name := range roles {
		var role models.Role

		err := db.Where("name = ?", name).First(&role).Error
		if err == nil {
			continue
		}

		db.Create(&models.Role{
			BaseModel: models.BaseModel{
				ID: uuid.New(),
			},
			Name: name,
		})
	}
}

func seedRolesAndPermissions(db *gorm.DB) {
	var adminRole models.Role
	db.Where("name = ?", "admin").First(&adminRole)

	var permissions []models.Permission
	db.Where("name IN ?", []string{
		"user:create",
		"user:view",
		"user:delete",
	}).Find(&permissions)

	if len(permissions) == 0 {
		return
	}

	// Evita duplicar associações
	db.Model(&adminRole).Association("Permissions").Replace(&permissions)
}

func seedAdminUser(db *gorm.DB) {
	var admin models.User
	if err := db.Where("email = ?", "admin@admin.com").First(&admin).Error; err == nil {
		log.Println("👤 Admin already exists")
		return
	}

	var company models.Company
	db.First(&company)

	var adminRole models.Role
	db.Where("name = ?", "admin").First(&adminRole)

	hashedPassword, _ := security.HashPassword("1234")

	admin = models.User{
		BaseModel: models.BaseModel{
			ID: uuid.New(),
		},
		Name:      "Admin",
		Email:     "admin@admin.com",
		Password:  hashedPassword,
		CompanyID: company.ID,
		RoleID:    adminRole.ID,
	}

	db.Create(&admin)
}
