package user

import (
	model "github.com/muharamdani/go-base-rest-api/core/models/user"
	request "github.com/muharamdani/go-base-rest-api/core/requests/user"
	"gorm.io/gorm"
)

func FetchAll(db *gorm.DB, out *[]model.Users, payload *request.FetchAllUser) error {
	return db.Limit(payload.Limit).Find(&out).Error
}