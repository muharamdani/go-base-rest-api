package user

import (
	"github.com/muharamdani/go-base-rest-api/core/models"
	"gorm.io/gorm"
)

func FetchAll(db *gorm.DB, out *[]models.Users, payload interface{}) error {
	return db.Find(&out).Error
}