package config

import (
	"github.com/joho/godotenv"
	"github.com/muharamdani/go-base-rest-api/utils"
	"log"
	"os"
)

var ENV *string

// Env Load environment variable from .env file
func Env(key string) string {
	rootPath := utils.GetRootPath()
	err := godotenv.Load(rootPath + (*ENV))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
