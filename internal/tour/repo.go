package tour

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repo interface {
	Insert(context.Context, Tour) int64
}

var _ repo = (*PgRepo)(nil)

type PgRepo struct {
	pool *pgxpool.Pool
}

func NewPgRepo(pool *pgxpool.Pool) PgRepo {
	return PgRepo{
		pool: pool,
	}
}

func (r PgRepo) Insert(ctx context.Context, tour Tour) int64 {
	var id int64
	err := r.pool.QueryRow(ctx, "insert into tour(title, description) values($1, $2) returning id", tour.Title, tour.Desc).Scan(&id)
	if err != nil {
		log.Println(err)
	}

	return id

}
