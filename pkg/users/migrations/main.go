package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

var migrations []*gormigrate.Migration
var m *gormigrate.Gormigrate

func init() {
	// Call migrate function for each migration here
	usersMigration := createUsersTable()
	migrations = append(migrations, usersMigration)
}

// Migrate will execute users migrations
func Migrate(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Migrated users table")
}

// RollbackMigration will rollback users migrations
func RollbackMigration(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
	}
	fmt.Println("Rolled back users table")
}
