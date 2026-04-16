package schedule

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

var _ repo = (*mockRepo)(nil)

func (m *mockRepo) Insert(ctx context.Context, schedule Schedule) int64 {
	args := m.Called(ctx, schedule)
	return int64(args.Int(0))
}

// tests that service calls the `Insert` method on repo
func TestInsertScheduleService(t *testing.T) {
	mockRepo := new(mockRepo)
	mockRepo.On("Insert", mock.Anything, mock.Anything).Return(1)

	scheduleService := NewScheduleService(mockRepo)
	newSchedule := Schedule{TourId: 1, StartsAt: time.Now().UTC().Round(time.Second).AddDate(0, 0, 15)}
	scheduleService.Insert(t.Context(), newSchedule)

	mockRepo.AssertExpectations(t)
}
