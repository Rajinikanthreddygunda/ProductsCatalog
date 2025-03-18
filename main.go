package main

import (
	products "apipro/Products"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:Manhattan99@tcp(127.0.0.1:3306)/productscatalog?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("database connection error:", err)
	}
	fmt.Println("Database connection successful")
	products.RegisterRoutes(db)
	fmt.Println("Sent request to routes")

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
