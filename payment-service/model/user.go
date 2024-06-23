package model

import (
	"time"

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
	Id             int  `gorm:"primarykey;index"`
	DisplayName    string
	Email          string
	Password       string
	PhoneNumber    string
	ProfilePicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	
	FromPayment []Payment `gorm:"foreignKey:FromUserId"` 
	ToPayment []Payment `gorm:"foreignKey:ToUserId"` 

}
