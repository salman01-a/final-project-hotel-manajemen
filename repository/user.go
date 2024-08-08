package repository

import (
	"database/sql"
	"errors"
	"final-project-sanbercode/models"
)

func Login(db *sql.DB, user *models.User) error {
	sqlStatement := "SELECT username, password, email, role FROM users WHERE username=$1 AND password=$2"
	row := db.QueryRow(sqlStatement, user.Username, user.Password)

	var dbUser models.User
	err := row.Scan(&dbUser.Username, &dbUser.Password, &dbUser.Email, &dbUser.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid username or password")
		}
		return err
	}

	user.Email = dbUser.Email
	user.Role = dbUser.Role
	return nil
}

func Register(db *sql.DB, user *models.User) error {
	sqlStatement := "INSERT INTO users (username, password, email, role) VALUES ($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(sqlStatement, user.Username, user.Password, user.Email, user.Role).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
