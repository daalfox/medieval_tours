package tour

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Tour struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}

type TourWithId struct {
	Id   int64
	Tour Tour
}

func InsertTour(ctx context.Context, pool *pgxpool.Pool, tour Tour) int64 {
	var id int64
	err := pool.QueryRow(ctx, "insert into tour(title, description) values($1, $2) returning id", tour.Title, tour.Desc).Scan(&id)
	if err != nil {
		log.Println(err)
	}

	return id
}
