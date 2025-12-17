package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"api-backend-saas/internal/testhelpers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJWTAuth_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/protected", JWTAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestJWTAuth_Authorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/protected", JWTAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	token := testhelpers.GenerateTestJWT(testhelpers.TestJWTClaims{
		UserID:      1,
		Role:        "admin",
		CompanyID:   1,
		Permissions: []string{"user:create"},
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	testhelpers.AuthRequest(req, token)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
