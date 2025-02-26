//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config config.yaml api.yaml

package api

import (
	"context"

	order_db "server/order_db"

	"github.com/jackc/pgx/v5"
)

type Order = order_db.Order

type OrdersApi struct {
	db *order_db.Queries
}

func NewOrdersApi(db *pgx.Conn) OrdersApi {
	queries := order_db.New(db)
	return OrdersApi{
		db: queries,
	}
}

func (o *OrdersApi) GetOrders(ctx context.Context, request GetOrdersRequestObject) (GetOrdersResponseObject, error) {
	orders, err := o.db.GetAllOrders(ctx)
	if err != nil {
		return GetOrders500Response{}, err
	}

	return GetOrders200JSONResponse(orders), nil
}

func (o *OrdersApi) PostOrders(ctx context.Context, request PostOrdersRequestObject) (PostOrdersResponseObject, error) {
	err := o.db.InsertOrder(ctx, order_db.InsertOrderParams{
		Symbol:    request.Body.Symbol,
		Price:     request.Body.Price,
		Quantity:  request.Body.Quantity,
		OrderType: request.Body.OrderType,
	})
	if err != nil {
		return PostOrders500Response{}, err
	}

	return PostOrders201Response{}, nil
}

type AllOrders []Order

type PostOrdersJSONRequestBody struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Quantity  int32   `json:"quantity"`
	OrderType string  `json:"orderType"`
}
