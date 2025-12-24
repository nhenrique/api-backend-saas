package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nhenrique/api-backend-saas/internal/infra/persistence/gorm/models"
	"github.com/nhenrique/api-backend-saas/internal/security"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

type CreateUserInput struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required"`
	RoleID   uuid.UUID `json:"role_id" binding:"required"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyIDStr := c.GetString("company_id")
	companyID, err := uuid.Parse(companyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company_id"})
		return
	}

	hashedPassword, err := security.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao processar senha"})
		return
	}

	user := models.User{
		BaseModel: models.BaseModel{
			ID: uuid.New(),
		},
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashedPassword,
		RoleID:    input.RoleID,
		CompanyID: companyID,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao criar usuário",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID.String(),
		"name":  user.Name,
		"email": user.Email,
	})
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	companyIDStr := c.GetString("company_id")
	companyID, err := uuid.Parse(companyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company_id"})
		return
	}

	var users []models.User

	err = h.DB.
		Where("company_id = ?", companyID).
		Select("id, name, email, role_id, company_id, created_at").
		Find(&users).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao buscar usuários",
		})
		return
	}

	// Opcional: transformar UUID → string no response
	response := make([]gin.H, 0, len(users))
	for _, u := range users {
		response = append(response, gin.H{
			"id":         u.ID.String(),
			"name":       u.Name,
			"email":      u.Email,
			"role_id":    u.RoleID.String(),
			"company_id": u.CompanyID.String(),
			"created_at": u.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}
