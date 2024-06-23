package model

import (
	"gorm.io/gorm"
)

// type User struct {
// 	Id             int       `json:"id"`
// 	DisplayName    string    `json:"display_name"`
// 	Email          string    `json:"email"`
// 	Password       string    `json:"password"`
// 	PhoneNumber    string    `json:"phone_number"`
// 	ProfilePicture string    `json:"profile_picture"`
// 	CreatedAt      time.Time `json:"created_at"`
// 	UpdatedAt      time.Time `json:"updated_at"`
// }

type User struct {
	gorm.Model
	DisplayName    string `gorm:"not null"`
	Email          string `gorm:"unique"`
	Password       string
	PhoneNumber    string `gorm:"unique"`
	ProfilePicture string

	FromPayment []Payment `gorm:"foreignKey:FromUserId"`
	ToPayment   []Payment `gorm:"foreignKey:ToUserId"`
}
