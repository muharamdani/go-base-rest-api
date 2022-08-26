//go:build integration || users || pkg || migration || all

package tests

import (
	"os"
	"testing"

	conn "github.com/muharamdani/go-base-rest-api/db"
	"github.com/muharamdani/go-base-rest-api/db/migrations"
)

func init() {
	if err := os.Setenv("APP_DIR", "go-base-rest-api"); err != nil {
		panic(err)
	}
	conn.Connect("testing")
}

// TestMigrate migrate users table
func TestMigrate(t *testing.T) {
	migrations.ExecuteMigrations(conn.DB)
}

// TestRollback rollback users table
func TestRollback(t *testing.T) {
	migrations.RollbackMigrations(conn.DB)
}
