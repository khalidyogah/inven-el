package repository

import (
	"database/sql"
	"inven-el/structs"
)

func GetAllTypes(db *sql.DB) (err error, results []structs.ItemType) {
	sql := "SELECT * FROM item_type"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var items = structs.ItemType{}

		err = rows.Scan(&items.Id, &items.Type, &items.CreatedAt, &items.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, items)
	}

	return
}

func GetItemType(db *sql.DB, id string) (err error, results []structs.ItemType) {
	sql := "SELECT * FROM item_type WHERE id = $1"

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var item = structs.ItemType{}

		err = rows.Scan(
			&item.Id,
			&item.Type,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}

		results = append(results, item)
	}

	return
}

func InsertItemType(db *sql.DB, item structs.ItemType) (err error) {
	sql := "INSERT INTO item_type (type) VALUES ($1)"

	errs := db.QueryRow(sql, item.Type)

	return errs.Err()
}

func UpdateItemType(db *sql.DB, item structs.ItemType) (err error) {
	sql := "UPDATE item_type SET type = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2"

	errs := db.QueryRow(sql, item.Type, item.Id)

	return errs.Err()
}

func DeleteItemType(db *sql.DB, item structs.ItemType) (err error) {
	sql := "DELETE FROM item_type WHERE id = $1"

	errs := db.QueryRow(sql, item.Id)

	return errs.Err()
}
