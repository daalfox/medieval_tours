package schedule

import (
	"context"
)

type ScheduleService struct {
	repo repo
}

func NewScheduleService(r repo) ScheduleService {
	return ScheduleService{
		repo: r,
	}
}

func (t ScheduleService) List(ctx context.Context) ([]ScheduleWithId, error) {
	return t.repo.List(ctx)
}
func (t ScheduleService) Insert(ctx context.Context, schedule Schedule) int64 {
	return t.repo.Insert(ctx, schedule)
}
