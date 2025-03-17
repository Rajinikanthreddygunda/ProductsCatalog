package products

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	h := NewHandler(db)
	http.HandleFunc("/products", h.CreateHandler(db))
	http.HandleFunc("/products/", h.DeleteHandler(db))
	http.HandleFunc("/products/update", h.UpdateHandler(db))
}
