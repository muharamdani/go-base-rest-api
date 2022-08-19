//go:build unit || users || pkg || all

package tests

import (
	"github.com/muharamdani/go-base-rest-api/pkg/users/models"
	"github.com/muharamdani/go-base-rest-api/pkg/users/repositories"
	"gorm.io/gorm"
	"testing"
)

// TestFetchByID test fetch user by id, example only, need to change using testify later
func TestFetchByID(t *testing.T) {
	type args struct {
		db  *gorm.DB
		out *models.Users
		id  int
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repositories.FetchByID(tt.args.db, tt.args.out, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("FetchByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
