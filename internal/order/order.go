package order

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Order struct {
	Desc string `json:"description"`
}

type OrderWithId struct {
	Id    int64
	Order Order
}

func InsertOrder(ctx context.Context, pool *pgxpool.Pool, order Order) int64 {
	var id int64
	err := pool.QueryRow(ctx, "insert into api_order(description) values ($1) returning id", order.Desc).Scan(&id)
	if err != nil {
		log.Println(err)
	}

	return id
}
