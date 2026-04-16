package tour

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"

	"github.com/daalfox/medieval_tours/internal/testhelpers"
)

func TestSaveOrder(t *testing.T) {
	connString, close := testhelpers.GetConnString(t)
	defer close()

	testhelpers.LoadTestData(t, connString)

	connPool, err := pgxpool.New(t.Context(), connString)
	if err != nil {
		t.Fatal(err)
	}

	newOrder := Tour{Desc: "some description"}
	id := InsertTour(t.Context(), connPool, newOrder)

	var desc string
	connPool.QueryRow(t.Context(), "select description from tour where id = $1", id).Scan(&desc)
	assert.Equal(t, desc, newOrder.Desc)
}
