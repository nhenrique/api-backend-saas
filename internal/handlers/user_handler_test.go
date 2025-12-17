package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"api-backend-saas/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	api := r.Group("/api")
	api.Use(middlewares.JWTAuth())

	api.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	req, _ := http.NewRequest("POST", "/api/users", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}
