package cache

import (
	"sort"
	"sync"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(100, 5*time.Minute)
	
	if cache == nil {
		t.Fatal("NewCache() returned nil")
	}
	
	if cache.Size() != 0 {
		t.Errorf("New cache should be empty, got size %d", cache.Size())
	}
}

func TestSetAndGet(t *testing.T) {
	cache := NewCache(100, 5*time.Minute)
	
	// Set a value
	cache.Set("key1", "value1", 0) // Use default TTL
	
	// Get the value
	value, found := cache.Get("key1")
	if !found {
		t.Error("Get() should find key1")
	}
	
	if value != "value1" {
		t.Errorf("Get() = %v, want %v", value, "value1")
	}
	
	// Get a non-existent key
	_, found = cache.Get("nonexistent")
	if found {
		t.Error("Get() should not find nonexistent key")
	}
}

func TestExpiration(t *testing.T) {
	cache := NewCache(100, 500*time.Millisecond)
	
	// Set a value with short TTL
	cache.Set("key1", "value1", 100*time.Millisecond)
	
	// Value should be available immediately
	_, found := cache.Get("key1")
	if !found {
		t.Error("Value should be available before expiration")
	}
	
	// Wait for expiration
	time.Sleep(200 * time.Millisecond)
	
	// Value should be gone
	_, found = cache.Get("key1")
	if found {
		t.Error("Value should be removed after expiration")
	}
	
	// Set a value with default TTL
	cache.Set("key2", "value2", 0)
	
	// Value should be available immediately
	_, found = cache.Get("key2")
	if !found {
		t.Error("Value with default TTL should be available before expiration")
	}
	
	// Wait for expiration
	time.Sleep(600 * time.Millisecond)
	
	// Value should be gone
	_, found = cache.Get("key2")
	if found {
		t.Error("Value with default TTL should be removed after expiration")
	}
}

func TestDelete(t *testing.T) {
	cache := NewCache(100, 5*time.Minute)
	
	// Set some values
	cache.Set("key1", "value1", 0)
	cache.Set("key2", "value2", 0)
	
	// Delete a value
	deleted := cache.Delete("key1")
	if !deleted {
		t.Error("Delete() should return true for existing key")
	}
	
	// Try to get the deleted value
	_, found := cache.Get("key1")
	if found {
		t.Error("Get() should not find deleted key")
	}
	
	// Other value should still be there
	_, found = cache.Get("key2")
	if !found {
		t.Error("Get() should still find key2")
	}
	
	// Delete non-existent key
	deleted = cache.Delete("nonexistent")
	if deleted {
		t.Error("Delete() should return false for non-existent key")
	}
}

func TestClear(t *testing.T) {
	cache := NewCache(100, 5*time.Minute)
	
	// Set some values
	cache.Set("key1", "value1", 0)
	cache.Set("key2", "value2", 0)
	
	// Clear the cache
	cache.Clear()
	
	// Cache should be empty
	if cache.Size() != 0 {
		t.Errorf("Cache should be empty after Clear(), got size %d", cache.Size())
	}
	
	// Should not find any keys
	_, found := cache.Get("key1")
	if found {
		t.Error("Get() should not find key after Clear()")
	}
}

func TestSize(t *testing.T) {
	cache := NewCache(100, 5*time.Minute)
	
	// Empty cache
	if cache.Size() != 0 {
		t.Errorf("Empty cache should have size 0, got %d", cache.Size())
	}
	
	// Add some items
	cache.Set("key1", "value1", 0)
	cache.Set("key2", "value2", 0)
	cache.Set("key3", "value3", 0)
	
	if cache.Size() != 3 {
		t.Errorf("Cache with 3 items should have size 3, got %d", cache.Size())
	}
	
	// Delete an item
	cache.Delete("key2")
	
	if cache.Size() != 2 {
		t.Errorf("Cache with 2 items should have size 2, got %d", cache.Size())
	}
	
	// Clear the cache
	cache.Clear()
	
	if cache.Size() != 0 {
		t.Errorf("Cleared cache should have size 0, got %d", cache.Size())
	}
}

func TestKeys(t *testing.T) {
	cache := NewCache(100, 5*time.Minute)
	
	// Empty cache
	keys := cache.Keys()
	if len(keys) != 0 {
		t.Errorf("Empty cache should have no keys, got %d", len(keys))
	}
	
	// Add some items
	cache.Set("key1", "value1", 0)
	cache.Set("key2", "value2", 0)
	cache.Set("key3", "value3", 0)
	
	// Get keys
	keys = cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Cache with 3 items should have 3 keys, got %d", len(keys))
	}
	
	// Sort keys for deterministic comparison
	sort.Strings(keys)
	expected := []string{"key1", "key2", "key3"}
	
	for i, key := range keys {
		if key != expected[i] {
			t.Errorf("keys[%d] = %s, want %s", i, key, expected[i])
		}
	}
}

func TestCapacity(t *testing.T) {
	// Create a cache with capacity 3
	cache := NewCache(3, 5*time.Minute)
	
	// Add items to fill the cache
	cache.Set("key1", "value1", 0)
	cache.Set("key2", "value2", 0)
	cache.Set("key3", "value3", 0)
	
	// All items should be in the cache
	for _, key := range []string{"key1", "key2", "key3"} {
		_, found := cache.Get(key)
		if !found {
			t.Errorf("Get(%q) should find value", key)
		}
	}
	
	// Access key1 to make it most recently used
	cache.Get("key1")
	
	// Add another item, should evict the least recently used (key2)
	cache.Set("key4", "value4", 0)
	
	// key2 should be evicted
	_, found := cache.Get("key2")
	if found {
		t.Error("key2 should be evicted (least recently used)")
	}
	
	// Other keys should still be there
	for _, key := range []string{"key1", "key3", "key4"} {
		_, found := cache.Get(key)
		if !found {
			t.Errorf("Get(%q) should find value", key)
		}
	}
}

func TestConcurrency(t *testing.T) {
	cache := NewCache(100, 5*time.Minute)
	
	// Number of goroutines
	n := 10
	
	// WaitGroup to wait for all goroutines
	var wg sync.WaitGroup
	wg.Add(n)
	
	// Start goroutines
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			
			// Each goroutine sets and gets its own keys
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				value := fmt.Sprintf("value-%d-%d", id, j)
				
				// Set value
				cache.Set(key, value, 0)
				
				// Get value
				val, found := cache.Get(key)
				if !found {
					t.Errorf("Get(%q) should find value", key)
				}
				
				if val != value {
					t.Errorf("Get(%q) = %v, want %v", key, val, value)
				}
				
				// Delete some keys
				if j%2 == 0 {
					cache.Delete(key)
				}
			}
		}(i)
	}
	
	// Wait for all goroutines to finish
	wg.Wait()
}

func TestPeriodicCleanup(t *testing.T) {
	// Create a cache with short cleanup interval
	cache := NewCache(100, 100*time.Millisecond)
	
	// Add items with short TTL
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Set(key, i, 50*time.Millisecond)
	}
	
	// Initially all items should be there
	if cache.Size() != 10 {
		t.Errorf("Cache should have 10 items, got %d", cache.Size())
	}
	
	// Wait for items to expire and cleanup to run
	time.Sleep(200 * time.Millisecond)
	
	// Cache should be empty after cleanup
	if cache.Size() != 0 {
		t.Errorf("Cache should be empty after cleanup, got size %d", cache.Size())
	}
}