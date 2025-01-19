package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data     map[string]cacheEntry
	mutex    *sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	cache := Cache{
		data:     make(map[string]cacheEntry),
		mutex:    &sync.RWMutex{},
		interval: interval}
	go func() {
		for range ticker.C {
			// fmt.Println("Clearing cache", interval)
			cache.ReadLoop()
		}
	}()
	return &cache
}

func (cache Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.data[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (cache Cache) Get(key string) ([]byte, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	entry, exists := cache.data[key]
	// fmt.Printf("Getting from cache: %s exists: %t\n", key, exists)

	return entry.val, exists
}

func (cache Cache) ReadLoop() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	for key, entry := range cache.data {
		if entry.createdAt.Before(time.Now().Add(-cache.interval)) {
			// fmt.Printf("Removing entry for: %s\n", key)
			delete(cache.data, key)
		}
	}
}
