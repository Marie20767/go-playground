package caching

import (
	"errors"
)

var (
	CacheMissError               = errors.New("cache miss")
	CacheEntryAlreadyExistsError = errors.New("entry with this key already exists in cache")
)

// A Cache that implements the 'Least Recently Used' eviction policy (i.e. if the cache is full and we try
// to add another item to it then the item that will be removed is the one that was used/read longest ago)
type LRUCache struct {
	capacity int

	// Stores all the keys of the items
	// Items at the end should be the most recently used and items at the start should be least recently used
	keys []string

	// Stores a map of keys to their values
	itemsMap map[string]string
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		keys:     []string{},
		itemsMap: map[string]string{},
	}
}

// If an item is in the cache:
// - Return it and move it to the top of the keys list (so it is considered most recently used)
// If an item is not in the cache:
// - Return a CacheMissError
func (c *LRUCache) Get(key string) (string, error) {
	// TODO: implement
}

// Remove an item from the cache (from itemsMap and from keys)
func (c *LRUCache) Remove(key string) {
	// TODO: implement
}

// If an item already exists in the cache, return an CacheEntryAlreadyExistsError
// If the cache is full:
// - Remove the least recently used item
// Add the new item to the cache
func (c *LRUCache) Add(key string, value string) error {
	// TODO: implement
}

// Returns the keys in the cache
func (c *LRUCache) Keys() []string {
	// TODO: implement
}
