package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password" `
	Email     string    `json:"email"`
	Role      string    `json:"role" `
	CreatedAt time.Time `json:"created_at"`
}
