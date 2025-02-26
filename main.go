package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"server/api"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	connectionString := "postgres://postgres:postgres@127.0.0.1:7777/postgres"
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		log.Fatalf("couldn't connect to database: %s", connectionString)
	}
	defer conn.Close(ctx)

	r := gin.Default()

	ordersApi := api.NewOrdersApi(conn)
	handlers := api.NewStrictHandler(&ordersApi, make([]api.StrictMiddlewareFunc, 0))
	api.RegisterHandlers(r, handlers)

	port := "8000"

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("localhost", port),
	}

	log.Fatal(s.ListenAndServe())
}
