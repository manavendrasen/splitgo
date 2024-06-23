package model

import "time"

type Payment struct {
	Id        int       `json:"id"`
	Amount    float32   `json:"amount"`
	From      User      `json:"from"`
	To        User      `json:"to"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
