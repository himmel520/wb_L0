package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type OrderPg struct {
	db *sql.DB
}

func NewOrderPg(db *sql.DB) *OrderPg {
	return &OrderPg{
		db: db,
	}
}

func (r *OrderPg) Create(order []byte) (string, error) {
	var id string
	err := r.db.QueryRow(`insert into orders (data) values ($1) returning id_orders`, order).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *OrderPg) GetByID(id string) ([]byte, error) {
	var order []byte
	err := r.db.QueryRow(`select (data) from orders where id_orders=($1)`, id).Scan(&order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
