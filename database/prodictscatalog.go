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
// GetProducts retrieves all products from the database.
func GetProducts(db *sql.DB) ([]model.Product, error) {
	query := "SELECT name, description, price FROM products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.Name, &product.Description, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
