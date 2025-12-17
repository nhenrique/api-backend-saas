package handlers

import (
	"net/http"

	"api-backend-saas/internal/models"
	"api-backend-saas/internal/security" // Certifique-se de ter esta função

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		RoleID   uint   `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash da senha antes de salvar no banco
	hashedPassword, err := security.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar a senha"})
		return
	}

	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashedPassword,
		RoleID:    input.RoleID,
		CompanyID: c.GetUint("company_id"),
	}

	h.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func (h *UserHandler) ListUsers(c *gin.Context) {

	companyID := c.GetUint("company_id")

	var users []models.User

	err := h.DB.
		Where("company_id = ?", companyID).
		Select("id, name, email, role_id, company_id, created_at").
		Find(&users).Error

	if err != nil {
		c.JSON(500, gin.H{
			"error": "erro ao buscar usuários",
		})
		return
	}

	c.JSON(200, users)
}
