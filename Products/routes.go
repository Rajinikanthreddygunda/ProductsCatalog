package products

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	h := NewHandler(db)
}
