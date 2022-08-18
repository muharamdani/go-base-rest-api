package repositories

import (
	"github.com/muharamdani/go-base-rest-api/pkg/users/models"
	"github.com/muharamdani/go-base-rest-api/pkg/users/requests"
	"gorm.io/gorm"
)

func FetchAll(db *gorm.DB, out *[]models.Users, payload *requests.GetUser) error {
	return db.Limit(payload.Limit).Offset(payload.Limit * (payload.Page - 1)).Find(&out).Error
}

func FetchByID(db *gorm.DB, out *models.Users, id int) error {
	return db.Where("id = ?", id).First(&out).Error
}

func Create(db *gorm.DB, out *models.Users) error {
	return db.Create(&out).Error
}

func Update(db *gorm.DB, id int, out *models.Users) error {
	return db.Model(models.Users{}).Where("id = ?", id).Updates(&out).Error
}

func Delete(db *gorm.DB, id int) error {
	return db.Where("id = ?", id).Delete(&models.Users{}).Error
}
