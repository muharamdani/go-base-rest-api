package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrations []*gormigrate.Migration
var m *gormigrate.Gormigrate

func ExecuteMigrations(db *gorm.DB) {
	// Call another migration here
	MigrateUsers(db)
}

func RollbackMigrations(db *gorm.DB) {
	// Call another rollback migration here
	RollbackUsers(db)
}
