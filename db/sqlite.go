package db

import (
	"github.com/muharamdani/go-base-rest-api/utils"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSqlite(dbName string) {
	rootPath := utils.GetRootPath()
	dbLocation := rootPath + dbName
	db, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		AllowGlobalUpdate:                        true,
	})
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	DB = db
}
