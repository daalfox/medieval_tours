package testhelpers

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func GetConnString(t testing.TB) (string, func()) {
	t.Helper()
	pgContainer, err := postgres.Run(t.Context(), "postgres:18",
		postgres.WithDatabase("medieval_tours"),
		postgres.BasicWaitStrategies(),
	)
	assert.NoError(t, err)
	deferFunc := func() {
		if err := testcontainers.TerminateContainer(pgContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}
	connString, err := pgContainer.ConnectionString(t.Context())
	assert.NoError(t, err)

	return connString, deferFunc
}

func LoadTestData(t testing.TB, connString string) {
	db, err := sql.Open("pgx", connString)
	assert.NoError(t, err)
	assert.NoError(t, goose.Up(db, "../../migrations"))
	testdata, err := os.ReadFile("../../testdata/test-data.sql")
	assert.NoError(t, err)
	db.ExecContext(t.Context(), string(testdata))
	db.Close()
}
