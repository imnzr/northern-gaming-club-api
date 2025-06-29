package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `gorm:"unique;not null" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	PhoneNumber string `json:"phone_number"`
}
