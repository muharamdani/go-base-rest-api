package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Users model
type Users struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	FirstName string    `json:"first_name" binding:"required" gorm:"type:varchar(100);not null"`
	LastName  string    `json:"last_name" binding:"required" gorm:"type:varchar(100);not null"`
	Username  string    `json:"username" binding:"required" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" binding:"required" gorm:"type:varchar(100);not null;unique"`
	Password  string    `json:"password" binding:"required" gorm:"type:varchar(50); not null"`
	// Time stamps
	gorm.Model
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New()
	return
}
