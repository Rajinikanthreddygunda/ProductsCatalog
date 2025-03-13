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

func (h Handler) GetProducts(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
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

// GetProducts retrieves all products from the database.
func GetProducts(db *sql.DB) ([]model.Product, error) {
	query := "SELECT name, description, price FROM products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
