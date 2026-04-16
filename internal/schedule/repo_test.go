package schedule

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"

	"github.com/daalfox/medieval_tours/internal/testhelpers"
)

// tests that repo correctly saves a new object
func TestSaveScheduleRepo(t *testing.T) {
	connString, close := testhelpers.GetConnString(t)
	defer close()

	testhelpers.LoadTestData(t, connString)

	connPool, err := pgxpool.New(t.Context(), connString)
	assert.NoError(t, err)

	scheduleRepo := NewPgRepo(connPool)

	newSchedule := Schedule{TourId: 1, StartsAt: time.Now().AddDate(0, 0, 15).Round(time.Second).UTC()}
	id := scheduleRepo.Insert(t.Context(), newSchedule)

	var schedule ScheduleWithId
	err = connPool.QueryRow(t.Context(), "select * from schedule where id = $1", id).Scan(&schedule.Id, &schedule.Schedule.TourId, &schedule.Schedule.StartsAt)
	assert.NoError(t, err)

	assert.Equal(t, schedule.Schedule.TourId, newSchedule.TourId)
	assert.WithinDuration(t, schedule.Schedule.StartsAt, newSchedule.StartsAt, time.Second)
}
