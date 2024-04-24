package repository

import (
	"database/sql"
	"inven-el/structs"
)

func GetAllItemConditions(db *sql.DB) (err error, results []structs.ItemCondition) {
	sql := "SELECT A.* FROM item_condition A " +
		"INNER JOIN (SELECT id, MAX(created_at) AS max_date FROM item_condition GROUP BY id) B " +
		"ON A.id = B.id AND A.created_at=B.max_date"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var item = structs.ItemCondition{}

		err = rows.Scan(
			&item.Id,
			&item.ItemId,
			&item.ImageUrl,
			&item.OverallConditon,
			&item.ImageUrl,
			&item.Status,
			&item.Note,
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

func GetItemCondition(db *sql.DB, id string) (err error, results []structs.ItemCondition) {
	sql := "SELECT A.* FROM item_condition A WHERE A.type_id = $1"

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var item = structs.ItemCondition{}

		err = rows.Scan(
			&item.Id,
			&item.ItemId,
			&item.ImageUrl,
			&item.OverallConditon,
			&item.ImageUrl,
			&item.Status,
			&item.Note,
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

func InsertItemCondition(db *sql.DB, item structs.ItemCondition) (err error) {
	sql := "INSERT INTO item ( item_id, image_url, lastcheck_date, overall_conditon, status, note) VALUES ($1, $2, (SELECT MAX(created_at) FROM item_condition B WHERE B.item_id = $1 ), $3, $4, $5)"

	errs := db.QueryRow(sql, item.ItemId, item.ImageUrl, item.OverallConditon, item.Status, item.Note)

	return errs.Err()
}

func UpdateItemCondition(db *sql.DB, item structs.ItemCondition) (err error) {
	sql := "UPDATE item SET item_id = $1, image_url = $2, overall_conditon = $3, status = $4, note = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $6"

	errs := db.QueryRow(sql, item.ItemId, item.ImageUrl, item.OverallConditon, item.Status, item.Note, item.Id)

	return errs.Err()
}

func DeleteItemCondition(db *sql.DB, item structs.ItemCondition) (err error) {
	sql := "DELETE FROM item WHERE id = $1"

	errs := db.QueryRow(sql, item.Id)

	return errs.Err()
}
