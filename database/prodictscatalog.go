package database

import (
	model "apipro/model"
	"database/sql"
)

// CreateBook inserts a new book record into the database.
func CreateBook(db *sql.DB, Product model.Product) error {
	// Insert the book record into the "books" table.
	query := "INSERT INTO products(name, description, price) VALUES (?, ?, ?)"
	_, err := db.Exec(query, Product.Name, Product.Description, Product.Price)
	if err != nil {
		return err
	}
	return nil
}
