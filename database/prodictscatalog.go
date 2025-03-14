package database

import (
	model "apipro/model"
	"database/sql"
	"fmt"
)

// CreateBook inserts a new book record into the database.
func CreateBook(db *sql.DB, Product model.Product) error {
	// Insert the book record into the "books" table.
	query := "INSERT INTO products(id,name, description, price) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, Product.ID, Product.Name, Product.Description, Product.Price)
	if err != nil {
		return err
	}
	return nil
}

// GetProductByID retrieves a product by its ID from the database.
func GetProductByID(db *sql.DB, id int64) (*model.Product, error) {
	var product model.Product

	query := "SELECT id, name, description, price FROM products WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with ID %d not found", id)
		}
		return nil, fmt.Errorf("error retrieving product: %w", err)
	}

	return &product, nil
}
