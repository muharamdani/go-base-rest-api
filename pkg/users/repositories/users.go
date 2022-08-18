package repositories

import (
	model "github.com/muharamdani/go-base-rest-api/pkg/users/models"
	request "github.com/muharamdani/go-base-rest-api/pkg/users/requests"
	"gorm.io/gorm"
)

func FetchAll(db *gorm.DB, out *[]model.Users, payload *request.FetchAllUser) error {
	return db.Limit(payload.Limit).Find(&out).Error
}
