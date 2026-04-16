package tour

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"

	"github.com/daalfox/medieval_tours/internal/testhelpers"
)

func TestSaveTour(t *testing.T) {
	connString, close := testhelpers.GetConnString(t)
	defer close()

	testhelpers.LoadTestData(t, connString)

	connPool, err := pgxpool.New(t.Context(), connString)
	if err != nil {
		t.Fatal(err)
	}

	newTour := Tour{Title: "some title", Desc: "some description"}
	id := InsertTour(t.Context(), connPool, newTour)

	var tour TourWithId
	connPool.QueryRow(t.Context(), "select * from tour where id = $1", id).Scan(&tour.Id, &tour.Tour.Title, &tour.Tour.Desc)
	assert.Equal(t, tour.Tour.Title, newTour.Title)
	assert.Equal(t, tour.Tour.Desc, newTour.Desc)
}
