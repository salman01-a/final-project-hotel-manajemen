package models

import "time"

type Payment struct {
	ID          int       `json:"id"`
	BookingID   int       `json:"booking_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
}
