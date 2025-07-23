package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// CacheItem represents an element stored in the cache
type CacheItem struct {
	Key       string
	Value     interface{}
	ExpiresAt time.Time
}

// Cache implements an in-memory cache system with expiration and limited capacity
type Cache struct {
	// TODO: Implement the necessary fields
}

// NewCache creates a new cache with the specified capacity and default TTL
func NewCache(capacity int, defaultTTL time.Duration) *Cache {
	// TODO: Implement this function
	return nil
}

// Set stores a value in the cache with a specific TTL
// If ttl is 0, the default TTL is used
// Returns true if added successfully, false if there was an error
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) bool {
	// TODO: Implement this function
	return false
}

// Get retrieves a value from the cache if it exists and hasn't expired
// Returns the value and a boolean indicating if it was found
func (c *Cache) Get(key string) (interface{}, bool) {
	// TODO: Implement this function
	return nil, false
}

// Delete removes a value from the cache
// Returns true if it was deleted, false if it didn't exist
func (c *Cache) Delete(key string) bool {
	// TODO: Implement this function
	return false
}

// Clear empties the cache
func (c *Cache) Clear() {
	// TODO: Implement this function
}

// Size returns the number of elements in the cache
func (c *Cache) Size() int {
	// TODO: Implement this function
	return 0
}

// Keys returns all active keys in the cache
func (c *Cache) Keys() []string {
	// TODO: Implement this function
	return nil
}