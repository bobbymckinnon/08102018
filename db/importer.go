package db

// Importer is used for all data importing actions
type Importer interface {
	// Import parses a csv file containing the list of providers and inserts them into the providers table
	Import() (string, error)
}
