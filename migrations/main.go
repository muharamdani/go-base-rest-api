package migrations

import (
	conn "github.com/muharamdani/go-base-rest-api/db"
	users "github.com/muharamdani/go-base-rest-api/pkg/users/migrations"
)

func ExecuteMigrations() {
	// Call another migration here
	conn.Connect("default")
	users.Migrate(conn.DB)
}
