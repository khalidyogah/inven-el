package repository

import (
	"database/sql"
	"inven-el/structs"
)

func GetAllItems(db *sql.DB) (err error, results []structs.ItemList) {
	sql := "SELECT * FROM item_list"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var items = structs.ItemList{}

		err = rows.Scan(&items.Id, &items.Name, &items.BoughtDate, &items.ImageUrl, &items.DepartmentId, &items.TypeId, &items.CreatedAt, &items.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, items)
	}

	return
}

func GetItem(db *sql.DB, id string) (err error, results []structs.ItemList) {
	sql := "SELECT * FROM item_list WHERE id = " + id

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var item = structs.ItemList{}

		err = rows.Scan(
			&item.Id,
			&item.Name,
			&item.BoughtDate,
			&item.ImageUrl,
			&item.DepartmentId,
			&item.TypeId,
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

func InsertItemList(db *sql.DB, item structs.ItemList) (err error) {
	sql := "INSERT INTO item_list (name, bought_date, image_url, department_id, type_id) VALUES ($1, $2, $3, $4, $5)"

	errs := db.QueryRow(sql, item.Name, item.BoughtDate, item.ImageUrl, item.DepartmentId, item.TypeId)

	return errs.Err()
}

func UpdateItemList(db *sql.DB, item structs.ItemList) (err error) {
	sql := "UPDATE item_list SET name = $1, bought_date=$2, image_url=$3, department_id=$4, type_id=$5, updated_at = CURRENT_TIMESTAMP WHERE id = $6"

	errs := db.QueryRow(sql, item.Name, item.BoughtDate, item.ImageUrl, item.DepartmentId, item.TypeId, item.Id)

	return errs.Err()
}

func DeleteItemList(db *sql.DB, item structs.ItemList) (err error) {
	sql := "DELETE FROM item_list WHERE id = $1"

	errs := db.QueryRow(sql, item.Id)

	return errs.Err()
}
