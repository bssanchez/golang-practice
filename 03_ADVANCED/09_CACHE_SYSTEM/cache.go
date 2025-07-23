package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// CacheItem representa un elemento almacenado en el cache
type CacheItem struct {
	Key       string
	Value     interface{}
	ExpiresAt time.Time
}

// Cache implementa un sistema de cache en memoria con expiración y capacidad limitada
type Cache struct {
	// TODO: Implementar los campos necesarios
}

// NewCache crea un nuevo cache con la capacidad y TTL por defecto especificados
func NewCache(capacity int, defaultTTL time.Duration) *Cache {
	// TODO: Implementar esta función
	return nil
}

// Set almacena un valor en el cache con una TTL específica
// Si ttl es 0, se usa el TTL por defecto
// Devuelve true si se añadió correctamente, false si hubo un error
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) bool {
	// TODO: Implementar esta función
	return false
}

// Get obtiene un valor del cache si existe y no ha expirado
// Devuelve el valor y un booleano indicando si se encontró
func (c *Cache) Get(key string) (interface{}, bool) {
	// TODO: Implementar esta función
	return nil, false
}

// Delete elimina un valor del cache
// Devuelve true si se eliminó, false si no existía
func (c *Cache) Delete(key string) bool {
	// TODO: Implementar esta función
	return false
}

// Clear vacía el cache
func (c *Cache) Clear() {
	// TODO: Implementar esta función
}

// Size devuelve el número de elementos en el cache
func (c *Cache) Size() int {
	// TODO: Implementar esta función
	return 0
}

// Keys devuelve todas las claves activas en el cache
func (c *Cache) Keys() []string {
	// TODO: Implementar esta función
	return nil
}