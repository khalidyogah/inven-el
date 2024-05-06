package repository

import (
	"database/sql"
	"inven-el/structs"

	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users (username, password) VALUES ($1,$2)"

	errs := db.QueryRow(sql, user.Username, user.Password)

	return errs.Err()
}

func Login(db *sql.DB, user structs.User) (err error) {
	var res structs.User
	sql := "SELECT * FROM users WHERE username = $1"

	row := db.QueryRow(sql, user.Username, user.Password)

	err = row.Scan(
		&res.Username,
		&res.Password,
	)
	if err != nil {
		return err
	}

	//decrypt + compare pass found
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(res.Password))
	if err != nil {
		return err
	}

	return row.Err()
}

func FindUser(db *sql.DB, username interface{}, user *structs.User) (err error) {
	sql := "SELECT * FROM users WHERE username = $1"

	row := db.QueryRow(sql, username)
	err = row.Scan(
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return err
	}

	return row.Err()
}
