package repository

import (
	"database/sql"
	"final-project-sanbercode/models"
)

func GetAllPayment(db *sql.DB) (result []models.Payment, err error) {
	sql := "SELECT * FROM payments"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var payment models.Payment

		err = rows.Scan(&payment.ID, &payment.BookingID, &payment.Amount, &payment.PaymentDate, &payment.Status)
		if err != nil {
			return
		}

		result = append(result, payment)
	}
	return
}

func InsertPayment(db *sql.DB, payment models.Payment) (err error) {
	sql := "INSERT INTO payments (booking_id, amount, payment_date, status) VALUES ($1, $2, $3, $4)"
	errs := db.QueryRow(sql, payment.BookingID, payment.Amount, payment.PaymentDate, payment.Status)
	return errs.Err()
}

func UpdatePayment(db *sql.DB, payment models.Payment) (err error) {
	sql := "UPDATE payments SET booking_id=$1, amount=$2, status=$3 WHERE id=$4"
	errs := db.QueryRow(sql, payment.BookingID, payment.Amount, payment.Status, payment.ID)
	return errs.Err()
}
func DeletePayment(db *sql.DB, payment models.Payment) (err error) {
	sql := "DELETE FROM payments WHERE id=$1"
	errs := db.QueryRow(sql, payment.ID)
	return errs.Err()
}
func GetPaymentByID(db *sql.DB, id int) (payment models.Payment, err error) {
	sql := "SELECT id, booking_id, amount, payment_date, status FROM payments WHERE id=$1"
	err = db.QueryRow(sql, id).Scan(&payment.ID, &payment.BookingID, &payment.Amount, &payment.PaymentDate, &payment.Status)
	return
}
