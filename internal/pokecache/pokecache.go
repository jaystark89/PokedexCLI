package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	pokeCache map[string]cacheEntry
	pokeMutex sync.Mutex
	interval  time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (cache *Cache) Add(key string, val []byte) {
	cache.pokeMutex.Lock()
	c := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.pokeCache[key] = c
	cache.pokeMutex.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.pokeMutex.Lock()
	value, ok := cache.pokeCache[key]
	v := value.val
	cache.pokeMutex.Unlock()
	return v, ok
}

func NewCache(interval time.Duration) *Cache {
	newMap := make(map[string]cacheEntry)
	c := Cache{
		pokeCache: newMap,
		interval:  interval,
	}
	go c.reapLoop()
	return &c
}

func (cache *Cache) reapLoop() {

	ticker := time.NewTicker(cache.interval)
	for range ticker.C {
		cache.pokeMutex.Lock()
		for c, e := range cache.pokeCache {
			if time.Since(e.createdAt) > cache.interval {
				delete(cache.pokeCache, c)
			}
		}
		cache.pokeMutex.Unlock()
	}
}
