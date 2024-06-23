package model

import "time"

type User struct {
	Id             int       `json:"id"`
	DisplayName    string    `json:"display_name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	PhoneNumber    string    `json:"phone_number"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}


