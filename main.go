package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/bobbymckinnon/08102018/db/importer"
	"github.com/bobbymckinnon/08102018/db/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	db          *sql.DB
	allowParams = []string{
		"state",
		"max_discharges",
		"min_discharges",
		"max_average_covered_charges",
		"min_average_covered_charges",
		"max_average_medicare_payments",
		"min_average_medicare_payments",
	}
)

// ImportProviders imports data from the providers.csv into the providers table
func ImportProviders(c *gin.Context) {
	// check ?importProviders is yes before starting import
	if "yes" != c.Request.URL.Query().Get("importProviders") {
		c.String(http.StatusOK, "the import has not been performed, required parameter missing")

		return
	}

	importer := importer.NewImporter(db)
	o, err := importer.Import()
	if err != nil {
		log.Printf("Error: %q", err)
		c.String(http.StatusInternalServerError, "An error has occurred")

		return
	}

	c.String(http.StatusOK, o)

	return
}

// GetProvidersByFilter retrieves providers from the postgres db
func GetProvidersByFilter(c *gin.Context) {
	err := ValidateGetProviders(c.Request.URL.Query())
	if err != nil {
		c.String(http.StatusOK, err.Error())

		return
	}

	store := postgres.NewStorage(db)
	o, err := store.GetProviders(c, c.Request.URL.Query())
	if err != nil {
		log.Printf("Error: %q", err)
		c.String(http.StatusInternalServerError, "not found")
	}

	c.String(http.StatusOK, o)
}

// ValidateGetProviders validates that at least one URL parameter has been provided,
// and that the URL parameter is allowed
func ValidateGetProviders(params map[string][]string) error {
	if len(params) <= 0 {
		return errors.New("Please provide at least one parameter")
	}
	ok := true
	for val, _ := range params {
		ok = false
		for _, t := range allowParams {
			if val == t {
				ok = true
			}
		}
	}
	if !ok {
		return errors.New("Parameter invalid")
	}

	return nil
}

// SetUpRouter defines routes and their associated functions
func SetUpRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "./static")

	// router for index
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	// router for importer
	router.GET("/importProviders", ImportProviders)

	// router for providers
	router.GET("/providers", GetProvidersByFilter)

	return router
}

func main() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL")) //+ " sslmode=disable"
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := SetUpRouter()
	router.Run(":" + port)
}
