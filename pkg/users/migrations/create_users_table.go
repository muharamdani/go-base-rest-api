package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/muharamdani/go-base-rest-api/pkg/users/models"
	"gorm.io/gorm"
)

func createUsersTable() *gormigrate.Migration {
	usersMigration := &gormigrate.Migration{
		ID: "201608301400",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Users{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&models.Users{})
		},
	}
	return usersMigration
}
