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
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
