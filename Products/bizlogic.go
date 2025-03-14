package api

import (
	"apipro/database"
	"apipro/model"
	"database/sql"
	"fmt"
)

type IBizlogic interface {
	CreateBookLogic(product model.Product) error
	GetProductLogic(id int64) (model.Product, error)
}

type Bizlogic struct {
	DB *sql.DB
}

func NewBizlogic(db *sql.DB) IBizlogic { // Return interface instead of struct pointer
	return &Bizlogic{DB: db}
}

func (bl *Bizlogic) CreateBookLogic(product model.Product) error {
	return database.CreateBook(bl.DB, product)
}

// GetProductLogic retrieves a product by ID.
func (bl *Bizlogic) GetProductLogic(id int64) (model.Product, error) {
	product, err := database.GetProductByID(bl.DB, id)
	if err != nil {
		return model.Product{}, fmt.Errorf("bizlogic: %w", err)
	}
	return *product, nil
}
