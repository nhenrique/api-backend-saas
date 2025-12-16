package config

import (
	"os"
	"time"
)

const JWTIssuer = "api-backend-saas"

var JWTSecret = []byte(getEnv("JWT_SECRET", "default-secret"))

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func JWTExpireDuration() time.Duration {
	exp := getEnv("JWT_EXPIRES_IN", "24h")
	d, _ := time.ParseDuration(exp)
	return d
}
