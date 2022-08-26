package seeders

import (
	"gorm.io/gorm"
	"log"

	"github.com/muharamdani/go-base-rest-api/pkg/users/models"

	"github.com/bxcodec/faker/v4"
)

func userSeeders(db *gorm.DB) error {
	var data []models.Users

	for i := 0; i < 20; i++ {
		data = append(data, models.Users{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Username:  faker.Username(),
			Email:     faker.Email(),
			Password:  faker.Password(),
		})
	}

	if db.Where("deleted_at IS NULL").Delete(&models.Users{}).Error != nil {
		return db.Error
	}

	if db.CreateInBatches(data, 20).Error != nil {
		return db.Error
	}

	log.Println("Users table seeded successfully")
	return nil
}
