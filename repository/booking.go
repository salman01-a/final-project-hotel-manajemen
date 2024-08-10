package repository

import (
	"database/sql"
	"final-project-sanbercode/models"
)

func GetAllBooking(db *sql.DB) (result []models.Booking, err error) {
	sql := "SELECT * FROM bookings"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var booking models.Booking

		err = rows.Scan(&booking.ID, &booking.RoomID, &booking.UserID, &booking.CheckIn, &booking.CheckOut, &booking.Status, &booking.CreatedAt)
		if err != nil {
			return
		}

		result = append(result, booking)
	}
	return
}

func InsertBooking(db *sql.DB, booking models.Booking) (err error) {
	sql := "INSERT INTO bookings (room_id, user_id, check_in, check_out, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	errs := db.QueryRow(sql, booking.RoomID, booking.UserID, booking.CheckIn, booking.CheckOut, booking.Status, booking.CreatedAt)
	return errs.Err()
}

func UpdateBooking(db *sql.DB, booking models.Booking) (err error) {
	sql := "`UPDATE bookings SET room_id=$1, user_id=$2, check_in=$3, check_out=$4, status=$5 WHERE id=$6"
	errs := db.QueryRow(sql, booking.RoomID, booking.UserID, booking.CheckIn, booking.CheckOut, booking.Status, booking.ID)
	return errs.Err()
}
func DeleteBooking(db *sql.DB, booking models.Booking) (err error) {
	sql := "DELETE FROM bookings WHERE id=$1"
	errs := db.QueryRow(sql, booking.ID)
	return errs.Err()
}
func GetBookingByID(db *sql.DB, id int) (booking models.Booking, err error) {
	sql := "SELECT id, room_id, user_id, check_in, check_out, status, created_at FROM bookings WHERE id=$1"
	err = db.QueryRow(sql, id).Scan(&booking.ID, &booking.RoomID, &booking.UserID, &booking.CheckIn, &booking.CheckOut, &booking.Status, &booking.CreatedAt)
	return
}
