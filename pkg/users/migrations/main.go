package migrations

import (
	"fmt"

	conn "github.com/muharamdani/go-base-rest-api/db"
	"github.com/muharamdani/go-base-rest-api/pkg/users/models"
)

func Migrate() {
	conn.Connect()
	conn.DB.Migrator().DropTable(&models.Users{})
	// Example another migrator command place below
	//conn.DB.Migrator().DropColumn(&models.Users{}, "address")
	fmt.Println("Migrate success")
}
