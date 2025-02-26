-- name: GetAllOrders :many
SELECT * FROM orders;

-- name: InsertOrder :exec
INSERT INTO orders(symbol, price, quantity, order_type) VALUES($1, $2, $3, $4);