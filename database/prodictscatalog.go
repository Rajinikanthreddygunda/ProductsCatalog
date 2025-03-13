package database

import (
	model "apipro/model"
	"database/sql"
)

// CreateProduct inserts a new product record into the database.
func CreateProduct(db *sql.DB, Product model.Product) error {
	// Insert the product record into the "products" table.
	query := "INSERT INTO products(id, name, description, price) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, Product.ID, Product.Name, Product.Description, Product.Price)
	if err != nil {
		return err
	}
	return nil
}
