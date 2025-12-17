package middlewares

import (
	"api-backend-saas/internal/database"
	"api-backend-saas/internal/models"

	"github.com/gin-gonic/gin"
)

func AuditLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		userID := c.GetUint("user_id")
		companyID := c.GetUint("company_id")

		log := models.AuditLog{
			UserID:    userID,
			CompanyID: companyID,
			Action:    c.FullPath(),
			Method:    c.Request.Method,
		}

		database.DB.Create(&log)
	}
}
