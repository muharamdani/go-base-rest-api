package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

var ENV *string

// Env Load environment variable from .env file
func Env(key string) string {
	projectDirName := os.Getenv("APP_DIR")
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	err := godotenv.Load(string(rootPath) + "/" + *ENV)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
