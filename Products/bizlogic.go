package Products

import (
	"apipro/database"
	"database/sql"
)

type IBizlogic interface {
	UpdateProductLogic(id int, updates map[string]interface{}) error
}

type Bizlogic struct {
	DB *sql.DB
}

func NewBizlogic(db *sql.DB) *Bizlogic {
	return &Bizlogic{DB: db}
}

func (bl *Bizlogic) UpdateProductLogic(id int, updates map[string]interface{}) error {
	return database.UpdateProduct(bl.DB, id, updates) // Now matches exactly
}
