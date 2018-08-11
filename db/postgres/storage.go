package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	xoproviders "github.com/bobbymckinnon/08102018/db/models"
	"github.com/gin-gonic/gin"
)

// Store implements the Store interface with Postgres database
type Store struct {
	pgdb *sql.DB
}

// Store returns a new postgres storage instance
func NewStorage(pgdb *sql.DB) *Store {
	return &Store{
		pgdb: pgdb,
	}
}

// Get
func (s *Store) Get(ID string) (string, error) {
	p, err := xoproviders.ProviderByID(s.pgdb, ID)
	if err != nil {
		log.Print(err)
		return "", err
	}

	pJson, _ := json.Marshal(p)
	if err != nil {
		log.Print(err)
		return "", err
	}

	return string(pJson), nil
}

// GetProviders filters the providers table based on our received query parameters
func (s *Store) GetProviders(c *gin.Context, params map[string][]string) (string, error) {
	sqlstring, args := s.BuildProvidersQuery(params)

	var providers []*xoproviders.Provider

	rows, err := s.pgdb.Query(sqlstring, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			provider = xoproviders.Provider{}
		)
		err := rows.Scan(
			&provider.ID,
			&provider.DrgDefinition,
			&provider.Name,
			&provider.StreetAddress,
			&provider.City,
			&provider.State,
			&provider.ZipCode,
			&provider.Hrrd,
			&provider.TotalDischarges,
			&provider.AverageCoveredCharges,
			&provider.AverageTotalPayments,
			&provider.AverageMedicarePayments,
		)
		if err != nil {
			return "", err
		}

		providers = append(providers, &provider)
	}

	pJson, err := json.Marshal(providers)
	if err != nil {
		log.Print(err)
		return "", err
	}

	return string(pJson), nil

}

// BuildProvidersQuery builds a dynamic sql query based on received query parameters
func (s *Store) BuildProvidersQuery(params map[string][]string) (string, []interface{}) {
	nextParamPlaceholder := getParamPlaceHolder()

	sqlstring := `SELECT
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
		`

	whereParts := make([]string, 0)
	args := make([]interface{}, 0)

	if val, ok := params["state"]; ok {
		// check that we are not allowing empty strings
		if len(val[0]) > 0 {
			whereParts = append(whereParts, "state = "+nextParamPlaceholder())
			args = append(args, strings.ToUpper(val[0]))
		}
	}

	if val, ok := params["max_discharges"]; ok {
		// check that we are only allowing ints in the query
		if _, err := strconv.Atoi(val[0]); err == nil {
			whereParts = append(whereParts, "total_discharges <= "+nextParamPlaceholder())
			args = append(args, val[0])
		}
	}

	if val, ok := params["min_discharges"]; ok {
		if _, err := strconv.Atoi(val[0]); err == nil {
			whereParts = append(whereParts, "total_discharges >= "+nextParamPlaceholder())
			args = append(args, val[0])
		}
	}

	if val, ok := params["max_average_covered_charges"]; ok {
		// check that we are only allowing floats in the query
		if _, err := strconv.ParseFloat(val[0], 64); err == nil {
			whereParts = append(whereParts, "average_covered_charges <= "+nextParamPlaceholder())
			args = append(args, val[0])
		}
	}

	if val, ok := params["min_average_covered_charges"]; ok {
		if _, err := strconv.ParseFloat(val[0], 64); err == nil {
			whereParts = append(whereParts, "average_covered_charges >= "+nextParamPlaceholder())
			args = append(args, val[0])
		}
	}

	if val, ok := params["max_average_medicare_payments"]; ok {
		if _, err := strconv.ParseFloat(val[0], 64); err == nil {
			whereParts = append(whereParts, "average_medicare_payments <= "+nextParamPlaceholder())
			args = append(args, val[0])
		}
	}

	if val, ok := params["min_average_medicare_payments"]; ok {
		if _, err := strconv.ParseFloat(val[0], 64); err == nil {
			whereParts = append(whereParts, "average_medicare_payments >= "+nextParamPlaceholder())
			args = append(args, val[0])
		}
	}

	if len(whereParts) > 0 {
		sqlstring += " WHERE " + strings.Join(whereParts, " AND ")
	}

	return sqlstring, args
}

func getParamPlaceHolder() func() string {
	paramPlaceHolderIndex := 1
	return func() string {
		result := fmt.Sprintf("$%d", paramPlaceHolderIndex)
		paramPlaceHolderIndex++
		return result
	}
}

func floatToString(num float32) string {
	return strconv.FormatFloat(float64(num), 'f', 2, 64)
}
