package importer

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	xoproviders "github.com/heroku/08102018/db/models"
)

const (
	// heroku does not allow more than 10,000 rows in a db
	maxDBRows = 199999
	// csv file containing providers
	providersCsv = "providers.csv"
)

// Importer implements the Importer interface with Postgres database
type Importer struct {
	pgdb *sql.DB
}

// NewImporter returns a new postgres storage instance
func NewImporter(pgdb *sql.DB) *Importer {
	return &Importer{
		pgdb: pgdb,
	}
}

// Import parses a csv file containing the list of providers and inserts them into the providers table
func (i *Importer) Import() (string, error) {
	var cnt int32

	csvFile, _ := os.Open(providersCsv)
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		td, err := StringToInt(line[8])
		if err != nil {
			continue
		}
		acc, err := StringToFloat(line[9])
		if err != nil {
			continue
		}
		atp, err := StringToFloat(line[10])
		if err != nil {
			continue
		}
		amp, err := StringToFloat(line[11])
		if err != nil {
			continue
		}

		provider := &xoproviders.Provider{
			DrgDefinition:           line[0],
			ID:                      line[1],
			Name:                    line[2],
			StreetAddress:           line[3],
			City:                    line[4],
			State:                   line[5],
			ZipCode:                 line[6],
			Hrrd:                    line[7],
			TotalDischarges:         td,
			AverageCoveredCharges:   acc,
			AverageTotalPayments:    atp,
			AverageMedicarePayments: amp,
		}

		err = provider.Insert(i.pgdb)
		if err != nil {
			log.Print(err)
		}

		cnt++

		// heroku does not allow more than 100,000 rows in a db
		if cnt == maxDBRows {
			break
		}
	}

	return "Inserted providers: " + strconv.Itoa(int(cnt)), nil
}

// StringToInt converts a string to an int16
func StringToInt(str string) (int16, error) {
	v, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return int16(v), nil
}

// StringToFloat converts a string to a float64
func StringToFloat(str string) (float64, error) {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}
