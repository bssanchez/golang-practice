# Exercise 9: Cache System

## Description
Implement an in-memory cache system with item expiration and size limit.

## Requirements
1. Implement the following structures and functions in the `cache.go` file:
   - Struct `CacheItem` to store a value with its expiration time
   - Struct `Cache` with fields to store items, maximum capacity, and default TTL
   - Method `NewCache(capacity int, defaultTTL time.Duration) *Cache` - Constructor
   - Method `Set(key string, value interface{}, ttl time.Duration) bool` - Stores a value with custom TTL
   - Method `Get(key string) (interface{}, bool)` - Gets a value if it exists and hasn't expired
   - Method `Delete(key string) bool` - Removes a value from the cache
   - Method `Clear()` - Empties the cache
   - Method `Size() int` - Returns the number of elements in the cache
   - Method `Keys() []string` - Returns all active keys

2. Considerations:
   - The cache must be thread-safe (use sync.Mutex or sync.RWMutex)
   - Implement an LRU (Least Recently Used) replacement policy when the cache is full
   - Expired elements must be automatically removed
   - Implement a periodic cleaning mechanism for expired elements

## Tests
Run `go test` to verify your implementation.