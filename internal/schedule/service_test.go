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
func (m *mockRepo) List(ctx context.Context) ([]ScheduleWithId, error) {
	args := m.Called(ctx)
	return args.Get(0).([]ScheduleWithId), args.Error(1)
}

// tests that service calls the `List` method on repo
func TestListScheduleService(t *testing.T) {
	mockRepo := new(mockRepo)
	mockRepo.On("List", mock.Anything).Return([]ScheduleWithId{}, nil)

	scheduleService := NewScheduleService(mockRepo)
	scheduleService.List(t.Context())

	mockRepo.AssertExpectations(t)
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
