package migrations

import (
	conn "github.com/muharamdani/go-base-rest-api/db"
	users "github.com/muharamdani/go-base-rest-api/pkg/users/migrations"
)

func init() {
	conn.Connect("default")
}

func ExecuteMigrations() {
	// Call another migration here
	users.Migrate(conn.DB)
}

func RollbackMigrations() {
	// Call another rollback migration here
	users.RollbackMigration(conn.DB)
}
