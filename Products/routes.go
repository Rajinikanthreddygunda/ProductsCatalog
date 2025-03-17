package Products

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	h := NewHandler(db)
	http.HandleFunc("/products/update", h.UpdateHandler()) // PUT - Update Product
}
