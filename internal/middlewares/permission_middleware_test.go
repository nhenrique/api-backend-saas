package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"api-backend-saas/internal/testhelpers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequirePermission_Forbidden(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET(
		"/secure",
		JWTAuth(),
		RequirePermission("user:create"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{})
		},
	)

	token := testhelpers.GenerateTestJWT(testhelpers.TestJWTClaims{
		UserID:      1,
		Permissions: []string{"user:view"},
	})

	req, _ := http.NewRequest("GET", "/secure", nil)
	testhelpers.AuthRequest(req, token)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusForbidden, resp.Code)
}

func TestRequirePermission_Allowed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET(
		"/secure",
		JWTAuth(),
		RequirePermission("user:create"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{})
		},
	)

	token := testhelpers.GenerateTestJWT(testhelpers.TestJWTClaims{
		UserID:      1,
		Permissions: []string{"user:create"},
	})

	req, _ := http.NewRequest("GET", "/secure", nil)
	testhelpers.AuthRequest(req, token)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
