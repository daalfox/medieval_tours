package tour

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"

	"github.com/daalfox/medieval_tours/internal/testhelpers"
)

// tests that repo correctly saves a new object
func TestSaveTourRepo(t *testing.T) {
	connString, close := testhelpers.GetConnString(t)
	defer close()

	testhelpers.LoadTestData(t, connString)

	connPool, err := pgxpool.New(t.Context(), connString)
	assert.NoError(t, err)

	tourRepo := NewPgRepo(connPool)

	t.Run("lists data", func(t *testing.T) {
		tours, err := tourRepo.List(t.Context())
		assert.NoError(t, err)

		assert.Equal(t, len(tours), 1)
	})

	t.Run("saves data", func(t *testing.T) {
		newTour := Tour{Title: "some title", Desc: "some description"}
		id := tourRepo.Insert(t.Context(), newTour)

		var tour TourWithId
		err = connPool.QueryRow(t.Context(), "select * from tour where id = $1", id).Scan(&tour.Id, &tour.Tour.Title, &tour.Tour.Desc)
		assert.NoError(t, err)

		assert.Equal(t, tour.Tour.Title, newTour.Title)
		assert.Equal(t, tour.Tour.Desc, newTour.Desc)
	})
}
