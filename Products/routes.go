package products

import (
	"database/sql"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	// Do nothing or return an error
	http.Error(w, "Route not implemented", http.StatusNotImplemented)
}

// func RegisterRoutes(db *sql.DB) {
// 	http.HandleFunc("/create", dummyHandler)
// 	http.HandleFunc("/books", dummyHandler)
// }

func RegisterRoutes(db *sql.DB) {
	//h := NewHandler(db)
	http.HandleFunc("/products", testHandler)

}
