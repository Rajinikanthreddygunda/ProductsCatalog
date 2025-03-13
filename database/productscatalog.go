package database

import (
	"database/sql"
)

// DeleteProduct removes a product from the database.
func DeleteProduct(db *sql.DB, productID int) error {
	// Delete the product record into the "products" table.
	query := "DELETE FROM products WHERE id = ?"
	_, err := db.Exec(query, productID)
	if err != nil {
		return err
	}
	return nil
}
