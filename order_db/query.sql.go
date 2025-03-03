// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package order_db

import (
	"context"
)

const getAllOrders = `-- name: GetAllOrders :many
SELECT id, symbol, price, quantity, order_type FROM orders
`

func (q *Queries) GetAllOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.Query(ctx, getAllOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.Symbol,
			&i.Price,
			&i.Quantity,
			&i.OrderType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertOrder = `-- name: InsertOrder :exec
INSERT INTO orders(symbol, price, quantity, order_type) VALUES($1, $2, $3, $4)
`

type InsertOrderParams struct {
	Symbol    string
	Price     float64
	Quantity  int32
	OrderType string
}

func (q *Queries) InsertOrder(ctx context.Context, arg InsertOrderParams) error {
	_, err := q.db.Exec(ctx, insertOrder,
		arg.Symbol,
		arg.Price,
		arg.Quantity,
		arg.OrderType,
	)
	return err
}
