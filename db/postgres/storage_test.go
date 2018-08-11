package postgres_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/bobbymckinnon/08102018/db/postgres"
	"github.com/stretchr/testify/assert"
)

var (
	db *sql.DB
)

// TestBuildProvidersQuery tests that our dynamic query string is built correctly
func TestBuildProvidersQuery(t *testing.T) {
	teststr := `SELECT
			id,
			drg_definition,
			name,
			street_address,
			city,
			state,
			zip_code,
			hrrd,
			total_discharges,
			average_covered_charges,
			average_total_payments,
			average_medicare_payments
		FROM
			public.providers
		 WHERE state = $1`

	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL")+" sslmode=disable")
	store := postgres.NewStorage(db)

	params := make(map[string][]string, 0)
	sqlstr, args := store.BuildProvidersQuery(params)

	assert.Empty(t, args)
	assert.NotNil(t, sqlstr)

	state := make([]string, 1)
	state[0] = "AL"
	params["state"] = state
	sqlstr, args = store.BuildProvidersQuery(params)

	assert.Equal(t, sqlstr, teststr)
}
