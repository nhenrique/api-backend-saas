package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"api-backend-saas/internal/services"
	"api-backend-saas/internal/testhelpers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin_InvalidEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := testhelpers.SetupTestDB()
	service := &services.AuthService{DB: db}
	handler := NewAuthHandler(service)

	r := gin.New()
	r.POST("/login", handler.Login)

	body := []byte(`{
		"email":"naoexiste@teste.com",
		"password":"123"
	}`)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestLogin_InvalidPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := testhelpers.SetupTestDB()

	testhelpers.CreateTestUser(db, "user@teste.com", "senha-correta")

	service := &services.AuthService{DB: db}
	handler := NewAuthHandler(service)

	r := gin.New()
	r.POST("/login", handler.Login)

	body := []byte(`{
		"email":"user@teste.com",
		"password":"senha-errada"
	}`)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestLogin_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := testhelpers.SetupTestDB()

	testhelpers.CreateTestUser(db, "user@teste.com", "senha123")

	service := &services.AuthService{DB: db}
	handler := NewAuthHandler(service)

	r := gin.New()
	r.POST("/login", handler.Login)

	body := []byte(`{
		"email":"user@teste.com",
		"password":"senha123"
	}`)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "token")
}
