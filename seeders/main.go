package seeders

import (
	"log"

	"github.com/muharamdani/go-base-rest-api/db"
	"gorm.io/gorm"
)

func Seed() {
	db.Connect()
	db.DB.Session(&gorm.Session{CreateBatchSize: 100, AllowGlobalUpdate: true})

	if err := userSeeders(db.DB); err != nil {
		panic(err)
	}

	log.Println("All table seeded successfully")
}