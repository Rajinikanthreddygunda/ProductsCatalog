package products

import (
	"apipro/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	biz IBizlogic
}

func NewHandler(db *sql.DB) Handler {
	return Handler{biz: NewBizlogic(db)}
}

func (h Handler) CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var product model.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if err := h.biz.CreateProductLogic(product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}
