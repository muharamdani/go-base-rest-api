//go:build unit || pkg || seeders || all

package tests

import (
	"os"
	"testing"

	conn "github.com/muharamdani/go-base-rest-api/db"
	"github.com/muharamdani/go-base-rest-api/db/seeders"
)

func init() {
	os.Setenv("APP_DIR", "go-base-rest-api")
	conn.Connect("testing")
}

func TestDBSeed(t *testing.T) {
	seeders.Seed(conn.DB)
}
