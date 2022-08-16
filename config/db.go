package config

import (
	"fmt"
	"github.com/muharamdani/gin-user-services/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// Connect creates a connection to postgresql database and
func Connect() {
	user := utils.Env("DB_USER")
	password := utils.Env("DB_PASS")
	host := utils.Env("DB_HOST")
	port := utils.Env("DB_PORT")
	database := utils.Env("DB_NAME")

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
