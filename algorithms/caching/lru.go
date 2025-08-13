package caching

import (
	"errors"
)

var (
	ErrCacheMiss               = errors.New("cache miss")
	ErrCacheEntryAlreadyExists = errors.New("entry with this key already exists in cache")
	ErrCacheCapacity           = errors.New("cache can't have 0 capacity")
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

func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity == 0 {
		return nil, ErrCacheCapacity
	}

	return &LRUCache{
		capacity: capacity,
		keys:     []string{},
		itemsMap: map[string]string{},
	}, nil
}

// If an item is in the cache:
// - Return it and move it to the top of the keys list (so it is considered most recently used)
// If an item is not in the cache:
// - Return a ErrCacheMiss
func (c *LRUCache) Get(key string) (string, error) {
	for i, cKey := range c.keys {
		if key == cKey {
			c.keys = append(c.keys[:i], c.keys[i+1:]...)
			c.keys = append(c.keys, cKey)
			return c.itemsMap[key], nil
		}
	}

	return "", ErrCacheMiss
}

// Remove an item from the cache (from itemsMap and from keys)
func (c *LRUCache) Remove(key string) {
	delete(c.itemsMap, key)
	c.capacity += 1

	for i, cKey := range c.keys {
		if cKey == key {
			c.keys = append(c.keys[:i], c.keys[i+1:]...)
			break
		}
	}
}

// If an item already exists in the cache, return an ErrCacheEntryAlreadyExists
// If the cache is full:
// - Remove the least recently used item
// Add the new item to the cache
func (c *LRUCache) Add(key string, value string) error {
	if _, exists := c.itemsMap[key]; exists {
		return ErrCacheEntryAlreadyExists
	}

	if len(c.keys) == c.capacity {
		c.keys = c.keys[1:]
	}

	c.keys = append(c.keys, key)
	c.itemsMap[key] = value

	return nil
}

// Returns the keys in the cache
func (c *LRUCache) Keys() []string {
	return c.keys
}
