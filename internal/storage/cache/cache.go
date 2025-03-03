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
	rowMu         sync.RWMutex
	short_to_long map[string]string
	long_to_short map[string]string
}

func New() *Cache {
	return &Cache{
		short_to_long: make(map[string]string),
		long_to_short: make(map[string]string),
	}
}

func (c *Cache) GetShortURL(long_url string) (string, error) {
	c.rowMu.Lock()
	defer c.rowMu.Unlock()

	if _, ok := c.long_to_short[long_url]; ok {
		return c.long_to_short[long_url], nil
	}
	return "", ErrURLNotFound

}

func (c *Cache) SaveURL(long_url, short_url string) (string, error) {

	existing_short_url, err := c.GetShortURL(long_url)
	if err == nil {
		return existing_short_url, nil
	}

	c.rowMu.Lock()
	{
		c.short_to_long[short_url] = long_url
		c.long_to_short[long_url] = short_url
	}
	c.rowMu.Unlock()

	return short_url, nil
}

func (c *Cache) GetURL(short_url string) (string, error) {

	//  Мьютекс ReadOnly не нужен, так как записи не меняются и не удаляются

	long_url, found := c.short_to_long[short_url]

	if !found {
		return "", ErrURLNotFound
	}

	return long_url, nil
}

func (c *Cache) GetCount() (int64, error) {
	return int64(len(c.short_to_long)), nil
}

func (c *Cache) Close() error {
	for k := range c.short_to_long {
		delete(c.short_to_long, k)
	}
	return nil
}
