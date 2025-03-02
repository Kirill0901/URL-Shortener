package postgresql

import (
	"flag"
	"testing"
)

var (
	host     = flag.String("host", "localhost", "address for postgresql")
	port     = flag.Int("port", 5432, "port for postgresql")
	user     = flag.String("user", "postgres", "user for postgresql")
	password = flag.String("password", "postgres", "password for postgresql")
	dbname   = flag.String("dbname", "postgres", "database name for postgresql")
)

func TestNewStorage(t *testing.T) {

	storage, err := New(*host, *port, *user, *password, *dbname)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer storage.Close()

	if storage == nil {
		t.Fatal("Expected storage to be non-nil")
	}

	if err := storage.Db.Ping(); err != nil {
		t.Fatalf("Database connection is not alive: %v", err)
	}

	err = storage.SaveURL("myFirstUrl", "myFirstShort")
	if err != nil {
		t.Errorf("Failed to save URL: %v", err)
	}

	long_url, err := storage.GetURL("myFirstShort")
	if err != nil || long_url != "myFirstUrl" {
		t.Errorf("Failed to get URL: %v", err)
	}
}
