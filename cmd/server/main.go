package main

import (
	"log"
	"os"

	"github.com/nhenrique/api-backend-saas/internal/database"
	"github.com/nhenrique/api-backend-saas/internal/routes"

	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "stage"
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Println("⚠️ Nenhum arquivo .env encontrado, usando env do sistema")
	}

	database.Connect()

	r := routes.SetupRouter()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
