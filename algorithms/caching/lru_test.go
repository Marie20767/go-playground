package caching_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/caching"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	t.Run("items can be added to and retrieved from the cache", func(t *testing.T) {
		cache, _ := caching.NewLRUCache(3)

		err := cache.Add("name", "Snowbell")
		assert.NoError(t, err)

		value, err := cache.Get("name")
		assert.NoError(t, err)
		assert.Equal(t, "Snowbell", value)
	})

	t.Run("returns a cache capacity error when trying to create an empty cache", func(t *testing.T) {
		_, err := caching.NewLRUCache(0)

		assert.ErrorIs(t, err, caching.ErrCacheCapacity)
	})

	t.Run("returns a cache miss error when trying to get an item that is not in the cache", func(t *testing.T) {
		cache, _ := caching.NewLRUCache(3)

		_, err := cache.Get("something")
		assert.ErrorIs(t, err, caching.ErrCacheMiss)
	})

	t.Run("items can be removed from the cache", func(t *testing.T) {
		cache, _ := caching.NewLRUCache(3)

		err := cache.Add("1", "lone")
		assert.NoError(t, err)

		value, err := cache.Get("1")
		assert.Equal(t, "", value)
		assert.ErrorIs(t, err, caching.ErrCacheMiss)
	})

	t.Run("if a key already exists in the cache, another item with the same key cannot be added", func(t *testing.T) {
		cache, _ := caching.NewLRUCache(3)

		err := cache.Add("7892", "bacon")
		assert.NoError(t, err)
		err = cache.Add("7892", "tomato")
		assert.ErrorIs(t, err, caching.ErrCacheEntryAlreadyExists)
	})

	t.Run("the least recently used/read item is removed from the cache when it is full", func(t *testing.T) {
		cache, _ := caching.NewLRUCache(3)

		// Add items to the cache until it is full
		err := cache.Add("1", "one")
		assert.NoError(t, err)
		err = cache.Add("2", "two")
		assert.NoError(t, err)
		err = cache.Add("3", "three")
		assert.NoError(t, err)

		assert.Equal(t, cache.Keys(), []string{"1", "2", "3"})

		// Simulate accessing stuff from the cache in a certain order
		_, err = cache.Get("2") // "2" is the first item we read so it will end up being the least recently used
		assert.NoError(t, err)
		_, err = cache.Get("1")
		assert.NoError(t, err)
		_, err = cache.Get("3")
		assert.NoError(t, err)

		// Assert order of keys is now in order of usage recency
		assert.Equal(t, []string{"2", "1", "3"}, cache.Keys())

		// Add a new item to the cache
		err = cache.Add("4", "four")
		assert.NoError(t, err)

		// Assert least recently used item was removed, and 4 added to end of the list
		assert.Equal(t, []string{"1", "3", "4"}, cache.Keys())

		// Assert that the 'Least Recently Used/Read' item is no longer in the cache
		_, err = cache.Get("2")
		assert.ErrorIs(t, err, caching.ErrCacheMiss)
	})
}
