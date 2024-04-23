package repository

import (
	"database/sql"
	"q3/structs"
)

func GetAllBooks(db *sql.DB) (err error, results []structs.Book) {
	sql := "SELECT * FROM book"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.CategoryId,
		)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := "INSERT INTO book (id, title, description, Image_url, Release_year, price, Total_page, thickness, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	errs := db.QueryRow(sql, book.Id, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryId)

	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := "UPDATE book SET title = $1, description = $2, Image_url = $3, Release_year = $4, price = $5, Total_page = $6, thickness = $7, category_id = $8, updated_at = CURRENT_TIMESTAMP WHERE id = $9"

	errs := db.QueryRow(sql, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryId, book.Id)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM book WHERE id = $1"

	errs := db.QueryRow(sql, book.Id)

	return errs.Err()
}
