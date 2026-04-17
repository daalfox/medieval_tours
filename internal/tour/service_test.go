package tour

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

var _ repo = (*mockRepo)(nil)

func (m *mockRepo) Insert(ctx context.Context, tour Tour) int64 {
	args := m.Called(ctx, tour)
	return int64(args.Int(0))
}

func (m *mockRepo) List(ctx context.Context) ([]TourWithId, error) {
	args := m.Called(ctx)
	return args.Get(0).([]TourWithId), args.Error(1)
}

// tests that service calls the `Insert` method on repo
func TestInsertTourService(t *testing.T) {
	mockRepo := new(mockRepo)
	mockRepo.On("Insert", mock.Anything, mock.Anything).Return(1)

	tourService := NewTourService(mockRepo)
	newTour := Tour{Title: "some title", Desc: "some description"}
	tourService.Insert(t.Context(), newTour)

	mockRepo.AssertExpectations(t)
}

// tests that service calls the `List` method on repo
func TestListTourService(t *testing.T) {
	mockRepo := new(mockRepo)
	mockRepo.On("List", mock.Anything).Return([]TourWithId{}, nil)

	tourService := NewTourService(mockRepo)
	tourService.List(t.Context())

	mockRepo.AssertExpectations(t)
}
