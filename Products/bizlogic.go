package products

import (
	"apipro/database"
	"apipro/model"
	"database/sql"
)

type IBizlogic interface {
	CreateProductLogic(product model.Product) error
	DeleteProductLogic(productID int) error
	UpdateProductLogic(product model.Product) error
}

type Bizlogic struct {
	DB *sql.DB
}

func NewBizlogic(db *sql.DB) *Bizlogic {
	return &Bizlogic{DB: db}
}

func (bl *Bizlogic) CreateProductLogic(product model.Product) error {
	return database.CreateProduct(bl.DB, product)
}

func (bl *Bizlogic) DeleteProductLogic(productID int) error {
	return database.DeleteProduct(bl.DB, productID)
}

func (bl *Bizlogic) UpdateProductLogic(updatedProduct model.Product) error {
	return database.UpdateProduct(bl.DB, updatedProduct)
}
