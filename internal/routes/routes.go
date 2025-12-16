package routes

import (
	"api-backend-saas/internal/database"
	"api-backend-saas/internal/handlers"
	"api-backend-saas/internal/middlewares"
	"api-backend-saas/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// AUTH
	authService := &services.AuthService{DB: database.DB}
	r.POST("/login", handlers.Login(authService))

	// PROTECTED API
	api := r.Group("/api")
	api.Use(middlewares.JWTAuth())

	// Admin-only
	api.GET("/admin/dashboard",
		middlewares.RequireRoles("admin"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "admin area"})
		},
	)

	// Manager + Admin
	api.GET("/reports",
		middlewares.RequireRoles("admin", "gerente"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "reports"})
		},
	)

	// Any authenticated user
	api.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user_id":    c.GetUint("user_id"),
			"role":       c.GetString("role"),
			"company_id": c.GetUint("company_id"),
		})
	})

	return r
}
