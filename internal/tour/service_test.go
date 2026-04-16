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

// tests that service calls the `Insert` method on repo
func TestInsertTourService(t *testing.T) {
	mockRepo := new(mockRepo)
	mockRepo.On("Insert", mock.Anything, mock.Anything).Return(1)

	tourService := NewTourService(mockRepo)
	newTour := Tour{Title: "some title", Desc: "some description"}
	tourService.Insert(t.Context(), newTour)

	mockRepo.AssertExpectations(t)
}
