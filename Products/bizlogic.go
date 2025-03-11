package api

import (
	"apipro/database"
	"apipro/model"
	"database/sql"
)

type IBizlogic interface {
	CreateBookLogic(product model.Product) error
}

type Bizlogic struct {
	DB *sql.DB
}

func NewBizlogic(db *sql.DB) *Bizlogic {
	return &Bizlogic{DB: db}
}

func (bl *Bizlogic) CreateBookLogic(product model.Product) error {
	return database.CreateBook(bl.DB, product)
}
