package repository

import (
	"database/sql"
	"inven-el/structs"
)

func CreateAccount(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO department (username, password) VALUES ($1,$2)"

	errs := db.QueryRow(sql, user.Username)

	return errs.Err()
}

func Login(db *sql.DB, user structs.User) (err error) {
	sql := "SELECT * FROM user WHERE username = $1 and password = $2"

	errs := db.QueryRow(sql, user.Username, user.Password)

	return errs.Err()
}
