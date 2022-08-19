//go:build integration || users || pkg || all || migrate

package tests

import (
	"testing"

	"github.com/muharamdani/go-base-rest-api/pkg/users/migrations"
	conn "github.com/muharamdani/go-base-rest-api/db"
)

// TestMigrate migrate users table, example only not real test case, need to change using testify later
func TestMigrate(t *testing.T) {
	//var tests []struct {
	//	name string
	//}
	conn.Connect("testing")

	migrations.Migrate(conn.DB)
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		migrations.Migrate()
	//	})
	//}
}
