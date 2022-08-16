package db

import (
	"fmt"
	"github.com/muharamdani/gin-user-services/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// Connect creates a connection to postgresql database and
func Connect() {
	user := config.Env("DB_USER")
	password := config.Env("DB_PASS")
	host := config.Env("DB_HOST")
	port := config.Env("DB_PORT")
	database := config.Env("DB_NAME")

	// postgres://user:password@host:port/database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	DB = db
}
