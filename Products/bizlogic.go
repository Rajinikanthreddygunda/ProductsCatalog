package products

import (
	"apipro/database"
	"apipro/model"
	"database/sql"
)

type IBizlogic interface {
	CreateProductLogic(product model.Product) error
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

