package migrations

import (
	"log"

	"github.com/muharamdani/go-base-rest-api/pkg/users/models"
	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	conn.AutoMigrate(&models.Users{})
	// Example another migrator command place below
	// conn.DB.Migrator().DropColumn(&models.Users{}, "address")
	log.Println("Migrate success")
}
