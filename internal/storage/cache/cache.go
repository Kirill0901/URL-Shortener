package cache

import (
	"errors"
	"sync"
)

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExists   = errors.New("URL exists")
)

type Cache struct {
	rowMu sync.RWMutex
	rows  map[string]string
}

func New() *Cache {
	return &Cache{
		rows: make(map[string]string),
	}
}

func (c *Cache) SaveURL(long_url, short_url string) error {
	c.rowMu.Lock()

	if _, ok := c.rows[long_url]; ok {
		return ErrURLExists
	}

	c.rows[long_url] = short_url

	c.rowMu.Unlock()
	return nil
}

func (c *Cache) GetURL(short_url string) (string, error) {

	//  Мьютекс ReadOnly не нужен, так как записи не меняются и не удаляются

	long_url, found := c.rows[short_url]

	if !found {
		return "", ErrURLNotFound
	}

	return long_url, nil
}

func (c *Cache) GetCount() (int64, error) {
	return int64(len(c.rows)), nil
}

func (c *Cache) Close() error {
	for k := range c.rows {
		delete(c.rows, k)
	}
	return nil
}
