//go:build integration || users || pkg || all

package tests

import (
	"github.com/muharamdani/go-base-rest-api/pkg/users/migrations"
	"testing"
)

// TestMigrate migrate users table, example only not real test case, need to change using testify later
func TestMigrate(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			migrations.Migrate()
		})
	}
}
