package tour

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Tour struct {
	Desc string `json:"description"`
}

type TourWithId struct {
	Id   int64
	Tour Tour
}

func InsertTour(ctx context.Context, pool *pgxpool.Pool, order Tour) int64 {
	var id int64
	err := pool.QueryRow(ctx, "insert into tour(description) values ($1) returning id", order.Desc).Scan(&id)
	if err != nil {
		log.Println(err)
	}

	return id
}
