package migrations

import (
	users "github.com/muharamdani/go-base-rest-api/pkg/users/migrations"
)

func ExecuteMigrations() {
	// Call another migration here
	users.Migrate()
}
