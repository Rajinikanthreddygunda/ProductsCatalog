package products

import (
	"apipro/database"
	"database/sql"
)

type IBizlogic interface {
	DeleteProductLogic(productID int) error
}

type Bizlogic struct {
	DB *sql.DB
}

func NewBizlogic(db *sql.DB) *Bizlogic {
	return &Bizlogic{DB: db}
}
func (bl *Bizlogic) DeleteProductLogic(productID int) error {
	return database.DeleteProduct(bl.DB, productID)
}
