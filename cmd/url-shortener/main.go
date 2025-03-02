package main

import (
	"errors"
	"flag"
	"log"
	"url-shortener/internal/storage/cache"
	"url-shortener/internal/storage/postgresql"
)

type Storage interface {
	SaveURL(long_url, short_url string) error
	GetURL(short_url string) (string, error)
	Close() error
}

func getStorage() (Storage, error) {
	storageType := flag.String("storage", "postgresql", "storage type: 'postgresql', 'cache'")
	host := flag.String("host", "localhost", "address for postgresql")
	port := flag.Int("port", 5432, "port for postgresql")
	user := flag.String("user", "postgres", "user for postgresql")
	password := flag.String("password", "postgres", "password for postgresql")
	dbname := flag.String("dbname", "postgres", "database name for postgresql")

	flag.Parse()

	switch *storageType {
	case "postgresql":
		return postgresql.New(*host, *port, *user, *password, *dbname)

	case "cache":
		return cache.New(), nil
	default:
		return nil, errors.New("unknown storage type")
	}
}

func main() {

	storage, err := getStorage()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer storage.Close()
}
