# Ejercicio 9: Sistema de Cache

## Descripción
Implementa un sistema de cache en memoria con expiración de elementos y límite de tamaño.

## Requisitos
1. Implementa las siguientes estructuras y funciones en el archivo `cache.go`:
   - Struct `CacheItem` para almacenar un valor con su tiempo de expiración
   - Struct `Cache` con campos para almacenar items, capacidad máxima y TTL por defecto
   - Método `NewCache(capacity int, defaultTTL time.Duration) *Cache` - Constructor
   - Método `Set(key string, value interface{}, ttl time.Duration) bool` - Almacena un valor con TTL personalizado
   - Método `Get(key string) (interface{}, bool)` - Obtiene un valor si existe y no ha expirado
   - Método `Delete(key string) bool` - Elimina un valor del cache
   - Método `Clear()` - Vacía el cache
   - Método `Size() int` - Devuelve el número de elementos en el cache
   - Método `Keys() []string` - Devuelve todas las claves activas

2. Consideraciones:
   - El cache debe ser thread-safe (utiliza sync.Mutex o sync.RWMutex)
   - Implementa una política de reemplazo LRU (Least Recently Used) cuando el cache está lleno
   - Los elementos expirados deben ser eliminados automáticamente
   - Implementa un mecanismo de limpieza periódica para elementos expirados

## Pruebas
Ejecuta `go test` para verificar tu implementación.