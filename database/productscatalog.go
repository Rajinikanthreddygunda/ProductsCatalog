package database

import (
	"database/sql"
	"fmt"
	"strings"
)

// // Update product by ID with dynamic fields
// func UpdateProduct(db *sql.DB, id int, updates map[string]interface{}) error {
// 	var setStatements []string
// 	var args []interface{}

// 	for key, value := range updates {
// 		setStatements = append(setStatements, key+" = ?")
// 		args = append(args, value)
// 	}

// 	query := "UPDATE products SET " + strings.Join(setStatements, ", ") + " WHERE ID = ?"
// 	args = append(args, id)

//		_, err := db.Exec(query, args...)
//		// return err
//	}
//
// Update product by ID with dynamic fields
func UpdateProduct(db *sql.DB, id int, updates map[string]interface{}) error {
	var setStatements []string
	var args []interface{}

	for key, value := range updates {
		setStatements = append(setStatements, fmt.Sprintf("%s = ?", key))
		args = append(args, value)
	}

	query := fmt.Sprintf("UPDATE products SET %s WHERE ID = ?", strings.Join(setStatements, ", "))
	args = append(args, id)

	_, err := db.Exec(query, args...)
	return err
}
