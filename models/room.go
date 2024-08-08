package models

import "time"

type Room struct {
	ID         int       `json:"id"`
	RoomNumber string    `json:"room_number"`
	Type       string    `json:"type"`
	Price      float64   `json:"price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
