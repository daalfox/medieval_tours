package tour

import (
	"context"
)

type TourService struct {
	repo repo
}

func NewTourService(r repo) TourService {
	return TourService{
		repo: r,
	}
}

func (t TourService) List(ctx context.Context) ([]TourWithId, error) {
	return t.repo.List(ctx)
}

func (t TourService) Insert(ctx context.Context, tour Tour) int64 {
	return t.repo.Insert(ctx, tour)
}
