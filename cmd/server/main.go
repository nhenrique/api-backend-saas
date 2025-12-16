package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"api-backend-saas/internal/database"
	"api-backend-saas/internal/routes"
)

func main() {
	database.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	log.Println("API running on :8080")
	r.Run(":8080")
}
