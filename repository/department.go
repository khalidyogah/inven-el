package repository

import (
	"database/sql"
	"inven-el/structs"
)

func GetAllDepartment(db *sql.DB) (err error, results []structs.Department) {
	sql := "SELECT * FROM department"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var departments = structs.Department{}

		err = rows.Scan(&departments.Id, &departments.Name, &departments.CreatedAt, &departments.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, departments)
	}

	return
}

func GetDepartment(db *sql.DB, id string) (err error, results []structs.Department) {
	sql := "SELECT * FROM department WHERE id = $1"

	defer errMsg()

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var department = structs.Department{}

		err = rows.Scan(
			&department.Id,
			&department.Name,
			&department.CreatedAt,
			&department.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}

		results = append(results, department)
	}

	return
}

func InsertDepartment(db *sql.DB, department structs.Department) (err error) {
	sql := "INSERT INTO department (name) VALUES ($1)"

	errs := db.QueryRow(sql, department.Name)

	return errs.Err()
}

func UpdateDepartment(db *sql.DB, department structs.Department) (err error) {
	sql := "UPDATE department SET name = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2"

	errs := db.QueryRow(sql, department.Name, department.Id)

	return errs.Err()
}

func DeleteDepartment(db *sql.DB, department structs.Department) (err error) {
	sql := "DELETE FROM department WHERE id = $1"

	errs := db.QueryRow(sql, department.Id)

	return errs.Err()
}
