package repository

import (
	"database/sql"
	"final-project-sanbercode/models"
)

func GetAllRoom(db *sql.DB) (result []models.Room, err error) {
	sql := "SELECT * FROM rooms"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var room models.Room

		err = rows.Scan(&room.ID, &room.RoomNumber, &room.Type, &room.Price, &room.Status, &room.CreatedAt)
		if err != nil {
			return
		}

		result = append(result, room)
	}
	return
}

func InsertRoom(db *sql.DB, room models.Room) (err error) {
	sql := "INSERT INTO rooms (room_number, type, price, status, created_at) VALUES ($1, $2, $3, $4, $5)"
	errs := db.QueryRow(sql, room.RoomNumber, room.Type, room.Price, room.Status, room.CreatedAt)
	return errs.Err()
}

func UpdateRoom(db *sql.DB, room models.Room) (err error) {
	sql := "UPDATE rooms SET room_number=$1, type=$2, price=$3, status=$4 WHERE id=$5"
	errs := db.QueryRow(sql, room.RoomNumber, room.Type, room.Price, room.Status, room.ID)
	return errs.Err()
}
func DeleteRoom(db *sql.DB, room models.Room) (err error) {
	sql := "DELETE FROM rooms WHERE id=$1"
	errs := db.QueryRow(sql, room.ID)
	return errs.Err()
}
func GetRoomByID(db *sql.DB, id int) (room models.Room, err error) {
	sql := "SELECT id, room_number, type, price, status, created_at FROM rooms WHERE id=$1"
	err = db.QueryRow(sql, id).Scan(&room.ID, &room.RoomNumber, &room.Type, &room.Price, &room.Status, &room.CreatedAt)
	return
}
