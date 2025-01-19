package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestGettingFromCache(t *testing.T) {
	cache := NewCache(5 * time.Second)
	entryKey := "first entry"
	entryVal := []byte("fake value")

	cache.Add(entryKey, entryVal)
	value, exists := cache.Get(entryKey)

	if !exists {
		t.Errorf("Problem with getting value %s from cache", entryKey)
	}
	if !bytes.Equal(value, entryVal) {
		t.Errorf("%v != %v", value, entryVal)
	}
}

func TestGettingFromCacheAfterCleanup(t *testing.T) {
	interval := 5 * time.Millisecond
	cache := NewCache(interval)
	entryKey := "first entry"
	entryVal := []byte("fake value")

	cache.Add(entryKey, entryVal)
	value, exists := cache.Get(entryKey)

	if !exists {
		t.Errorf("Problem with getting value: %s", entryKey)
	}
	if !bytes.Equal(value, entryVal) {
		t.Errorf("%v != %v", value, entryVal)
	}

	time.Sleep(2 * interval)

	value, exists = cache.Get(entryKey)

	if exists {
		t.Errorf("Problem with cleaning value: %s it still exists", entryKey)
	}
	if !bytes.Equal(value, nil) {
		t.Errorf("%v != %v", value, nil)
	}
}
