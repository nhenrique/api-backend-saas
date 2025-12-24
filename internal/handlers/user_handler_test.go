package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nhenrique/api-backend-saas/internal/infra/persistence/gorm/models"
	"github.com/nhenrique/api-backend-saas/internal/middlewares"
	"github.com/nhenrique/api-backend-saas/internal/testhelpers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := testhelpers.SetupTestDB()
	handler := NewUserHandler(db)

	r := gin.New()
	api := r.Group("/api")
	api.Use(middlewares.JWTAuth())
	api.POST("/users", handler.CreateUser)

	body := []byte(`{
		"name":"Teste",
		"email":"teste@teste.com",
		"password":"123",
		"role_id":1
	}`)

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestCreateUser_Forbidden(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := testhelpers.SetupTestDB()
	handler := NewUserHandler(db)

	r := gin.New()
	api := r.Group("/api")
	api.Use(
		middlewares.JWTAuth(),
		middlewares.RequirePermission("user:create"),
	)
	api.POST("/users", handler.CreateUser)

	token := testhelpers.GenerateTestJWT(testhelpers.TestJWTClaims{
		UserID:      1,
		CompanyID:   1,
		Permissions: []string{"user:view"},
	})

	body := []byte(`{
		"name":"Teste",
		"email":"teste@teste.com",
		"password":"123",
		"role_id":1
	}`)

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	testhelpers.AuthRequest(req, token)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusForbidden, resp.Code)
}

func TestCreateUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := testhelpers.SetupTestDB()

	// Seed mínimo necessário
	db.Create(&models.Company{Name: "Empresa"})
	db.Create(&models.Role{ID: 1, Name: "User"})

	handler := NewUserHandler(db)

	r := gin.New()
	api := r.Group("/api")
	api.Use(
		middlewares.JWTAuth(),
		middlewares.RequirePermission("user:create"),
	)
	api.POST("/users", handler.CreateUser)

	token := testhelpers.GenerateTestJWT(testhelpers.TestJWTClaims{
		UserID:      1,
		CompanyID:   1,
		Permissions: []string{"user:create"},
	})

	body := []byte(`{
		"name":"João",
		"email":"joao@teste.com",
		"password":"123456",
		"role_id":1
	}`)

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	testhelpers.AuthRequest(req, token)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}
