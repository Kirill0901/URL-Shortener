package cache

import "testing"

func TestCache(t *testing.T) {

	cache := New()

	defer cache.Close()

	cache.SaveURL("key1", "value1")
	cache.SaveURL("key2", "value2")

	item, err := cache.GetURL("key1")

	if err != nil {
		t.Errorf("Failed to retrieve item from cache")
	} else if item != "value1" {
		t.Errorf("Incorrect value retrieved from cache")
	}

	cnt, err := cache.GetCount()
	if err != nil || cnt != 2 {
		t.Errorf("Incorrect count of items in cache")
	}
}
