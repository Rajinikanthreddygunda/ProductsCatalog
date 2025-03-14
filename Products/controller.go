package products

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	biz IBizlogic
}

func NewHandler(db *sql.DB) Handler {
	return Handler{biz: NewBizlogic(db)}
}

func (h Handler) DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract product ID from URL
		idStr := strings.TrimPrefix(r.URL.Path, "/products/")
		productID, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		// Call business logic to delete product
		if err := h.biz.DeleteProductLogic(productID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Product with ID %d deleted successfully", productID)))
	}
}
