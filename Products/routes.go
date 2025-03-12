package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	h := NewHandler(db) //newline
	http.HandleFunc("/products", h.CreateHandler(db))
	http.HandleFunc("/products", h.GetHandler(db))
}
