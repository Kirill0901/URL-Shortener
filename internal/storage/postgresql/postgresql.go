package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	db       *sql.DB
}

func (s *Storage) execPrepare(script string) error {
	stmt, err := s.db.Prepare(script)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func New(host string, port int, username, password, database string) (*Storage, error) {
	s := &Storage{
		host:     host,
		port:     port,
		user:     username,
		password: password,
		dbname:   database,
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.host, s.port, s.user, s.password, s.dbname)

	var err error
	s.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = s.db.Ping(); err != nil {
		return nil, err
	}

	err = s.execPrepare(`
		CREATE TABLE IF NOT EXISTS url_shortener (
			id SERIAL PRIMARY KEY,
			long_url VARCHAR(255) NOT NULL UNIQUE,
			short_url VARCHAR(255) NOT NULL UNIQUE);
	`)
	if err != nil {
		return nil, err
	}

	err = s.execPrepare(`
		CREATE INDEX IF NOT EXISTS idx_short_url ON url_shortener(short_url);
	`)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL")
	return s, nil
}

func (s *Storage) SaveURL(long_url, short_url string) error {
	stmt, err := s.db.Prepare("INSERT INTO url_shortener (long_url, short_url) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(long_url, short_url); err != nil {
		return err
	}

	return err
}

func (s *Storage) GetURL(short_url string) (string, error) {
	stmt, err := s.db.Prepare("SELECT long_url FROM url_shortener WHERE short_url = $1")
	if err != nil {
		return "", err
	}

	var long_url string
	err = stmt.QueryRow(short_url).Scan(&long_url)
	if err != nil {
		return "", err
	}

	return long_url, nil

}

func (s *Storage) Close() error {
	return s.db.Close()
}
