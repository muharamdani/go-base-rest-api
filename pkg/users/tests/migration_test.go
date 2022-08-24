//go:build integration || users || pkg || migrate || all

package tests

import (
	"github.com/muharamdani/go-base-rest-api/pkg/users/migrations"
	"os"
	"testing"

	conn "github.com/muharamdani/go-base-rest-api/db"
)

// TestMigrate migrate users table, example only not real test case, need to change using testify later
func TestMigrate(t *testing.T) {
	os.Setenv("APP_DIR", "go-base-rest-api")
	conn.Connect("testing")

	migrations.Migrate(conn.DB)
}
