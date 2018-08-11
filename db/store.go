package db

// Store is used for all data actions on our postgres db
type Store interface {
	// Get retrieves a single provider from the providers table based on the ID
	Get(ID string) (string, error)
	// GetProviders retrieves a set of providers from the providers table based on received URL parameters
	GetProviders(params map[string][]string) (string, error)
	// BuildProvidersQuery builds a dynamic sql query based on received query parameters
	BuildProvidersQuery(params map[string][]string) (string, []interface{})
}
