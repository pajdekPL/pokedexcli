package pokecache

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestGetAdd(t *testing.T) {
	interval := 5 * time.Microsecond
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("Some Fake Data"),
		},
		{
			key: "https://anotherexample.com/path",
			val: []byte(""),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Execute case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)

			val, exists := cache.Get(c.key)
			if !exists {
				t.Errorf("Problem getting added value for: %s", c.key)
			}
			if !bytes.Equal(val, c.val) {
				t.Errorf("Corrupted value in cache: %v != %v", val, c.val)
			}
		})
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
		return
	}
	if !bytes.Equal(value, entryVal) {
		t.Errorf("%v != %v", value, entryVal)
		return
	}

	time.Sleep(2 * interval)

	value, exists = cache.Get(entryKey)

	if exists {
		t.Errorf("Problem with cleaning value: %s it still exists", entryKey)
		return
	}
	if !bytes.Equal(value, nil) {
		t.Errorf("%v != %v", value, nil)
		return
	}
}
