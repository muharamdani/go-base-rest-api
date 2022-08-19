package seeders

import (
	"log"

	"github.com/muharamdani/go-base-rest-api/pkg/users/models"

	"gorm.io/gorm"
)

func userSeeders(db *gorm.DB) error{
	data := []models.Users{
		{
			FirstName: "Admin",
			LastName: "Istrator",
			Username: "admin",
			PhoneNumber: "081234567890",
			Email: "admin@paninti.com",
			Address: "Jl. Paninti No.1",
		},
		{
			FirstName: "Customer",
			LastName: "1",
			Username: "customer1",
			PhoneNumber: "081234567890",
			Email: "cust@paninti.com",
			Address: "Jl. Paninti No.1",
		},
	}

	if db.Delete(&models.Users{}).Error != nil {
		return db.Error
	}

	if db.CreateInBatches(data, 100).Error != nil {
		return db.Error
	}

	log.Println("Users table seeded successfully")
	return nil
}