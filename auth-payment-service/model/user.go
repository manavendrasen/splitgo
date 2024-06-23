package model

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	DisplayName    string `gorm:"not null"`
	Email          string `gorm:"unique"`
	Password       string
	PhoneNumber    string `gorm:"unique"`
	ProfilePicture string

	FromPayment []Payment `gorm:"foreignKey:FromUserId"`
	// ToPayment   []Payment `gorm:"foreignKey:ToUserId"`
}
