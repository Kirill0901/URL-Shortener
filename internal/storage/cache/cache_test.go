package cache

import "testing"

func TestCache(t *testing.T) {

	cache := New()

	defer cache.Close()

	short1 := "short1"

	cache.SaveURL("long1", short1)
	cache.SaveURL("long2", "short2")

	short_url, err := cache.SaveURL("long1", "short3")
	if err != nil || short_url != short1 {
		t.Errorf("expected short1, got %s", short_url)
	}

	long_url, err := cache.GetURL("short1")

	if err != nil {
		t.Errorf("Failed to retrieve item from cache")
	} else if long_url != "long1" {
		t.Errorf("Incorrect value retrieved from cache")
	}

	_, err = cache.GetURL("short100")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	cnt, err := cache.GetCount()
	if err != nil || cnt != 2 {
		t.Errorf("Incorrect count of items in cache")
	}
}
