package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Env Load environment variable from .env file
func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
