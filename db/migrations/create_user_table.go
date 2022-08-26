package migrations

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Users struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	FirstName string    `json:"first_name" binding:"required" gorm:"type:varchar(100);not null"`
	LastName  string    `json:"last_name" binding:"required" gorm:"type:varchar(100);not null"`
	Username  string    `json:"username" binding:"required" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" binding:"required" gorm:"type:varchar(100);not null;unique"`
	Password  string    `json:"password" binding:"required" gorm:"type:varchar(50); not null"`
	gorm.Model
}

func init() {
	migrations = []*gormigrate.Migration{
		{
			ID: "20220826",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(Users{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(Users{})
			},
		},
	}
}

// MigrateUsers will execute users migrations
func MigrateUsers(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Migrated users table")
}

// RollbackUsers will rollback users migrations
func RollbackUsers(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
	}
	fmt.Println("Rolled back users table")
}
