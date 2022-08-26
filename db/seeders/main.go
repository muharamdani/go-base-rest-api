package seeders

import (
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Run seeders here
	db.Session(&gorm.Session{CreateBatchSize: 100, AllowGlobalUpdate: true})
	if err := userSeeders(db); err != nil {
		panic(err)
	}
}
