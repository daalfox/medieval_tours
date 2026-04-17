package schedule

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo interface {
	List(context.Context) ([]ScheduleWithId, error)
	Insert(context.Context, Schedule) int64
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

func (r PgRepo) List(ctx context.Context) ([]ScheduleWithId, error) {
	rows, _ := r.pool.Query(ctx, "select * from schedule")
	schedules, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (ScheduleWithId, error) {
		s := ScheduleWithId{}
		if err := row.Scan(&s.Id, &s.Schedule.TourId, &s.Schedule.StartsAt); err != nil {
			return s, err
		}
		return s, nil
	})
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r PgRepo) Insert(ctx context.Context, schedule Schedule) int64 {
	var id int64
	err := r.pool.QueryRow(ctx, "insert into schedule(tour_id, starts_at) values($1, $2) returning id", schedule.TourId, schedule.StartsAt).Scan(&id)
	if err != nil {
		log.Println(err)
	}

	return id

}
