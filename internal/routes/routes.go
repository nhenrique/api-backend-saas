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

	// ========================
	// Handlers
	// ========================
	authService := &services.AuthService{DB: database.DB}
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(database.DB)

	// ========================
	// Public routes
	// ========================
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/login", authHandler.Login)

	// ========================
	// Protected routes
	// ========================
	api := r.Group("/api")
	api.Use(
		middlewares.JWTAuth(),
		middlewares.EnforceTenant(),
		middlewares.AuditLog(),
	)

	api.POST(
		"/users",
		middlewares.RequirePermission("user:create"),
		userHandler.CreateUser,
	)

	api.GET(
		"/users",
		middlewares.RequirePermission("user:list"),
		userHandler.ListUsers,
	)

	return r
}
