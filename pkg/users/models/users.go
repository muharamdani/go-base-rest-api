package models

// Users model Example
type Users struct {
	Id          int    `json:"id" gorm:"primary_key;autoIncrement"`
	FirstName   string `json:"first_name" binding:"required" gorm:"string;not null"`
	LastName    string `json:"last_name" binding:"required" gorm:"string;not null"`
	Username    string `json:"username" binding:"required" gorm:"string;not null"`
	PhoneNumber string `json:"phone_number" binding:"required" gorm:"string;not null"`
	Email       string `json:"email" binding:"required" gorm:"string;not null"`
	Address     string `json:"address" binding:"required" gorm:"text; not null"`
}

type GetUserPayload struct {
	Limit int
}
