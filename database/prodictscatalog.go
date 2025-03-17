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

func DeleteProduct(db *sql.DB, productID int) error {
	// Delete the product record into the "products" table.
	query := "DELETE FROM products WHERE id = ?"
	_, err := db.Exec(query, productID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *sql.DB, updatedProduct model.Product) error {
	query := "UPDATE products SET description = ?, price = ? WHERE Name = ?"
	_, err := db.Exec(query, updatedProduct.Description, updatedProduct.Price)
	if err != nil {
		return err
	}

	return nil
}
