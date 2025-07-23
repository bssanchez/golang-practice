# Ejercicios de pŕactica GO - Estilo TDD

Este repositorio contiene 13 ejercicios de Go organizados por nivel de dificultad. Cada ejercicio tiene test cases pre-escritos que debes hacer pasar implementando las funciones correspondientes.

## Consejo personal

Este es un repositorio para practicar y aprender GO; la idea es que te apoyes de documentación oficial, foros, plataformas de Q&A tipo stackoverflow, etc., procura no dejar que un agrente de IA los resuelva del todo por ti, puedes apoyarte un poco pero la idea es que entiendas completamente lo que estás desarrollando.

## Estructura del Proyecto

```
./
├── BASIC/
│   ├── 01_CALCULATOR/
│   ├── 02_STRING_UTILS/
│   └── 03_ARRAY_OPERATIONS/
├── INTERMEDIATE/
│   ├── 04_BANK_ACCOUNT/
│   ├── 05_TODO_LIST/
│   ├── 06_FILE_PROCESSOR/
│   └── 07_HTTP_CLIENT/
├── ADVANCED/
│   ├── 08_CONCURRENT_DOWNLOADER/
│   ├── 09_CACHE_SYSTEM/
│   └── 10_WEB_SCRAPER/
└── IA/
    ├── 11_SENTIMENT_ANALYZER/
    ├── 12_IMAGE_CLASSIFIER/
    └── 13_CHATBOT_ENGINE/
```

## Instrucciones

1. Para cada ejercicio, navega al directorio correspondiente
2. Lee el archivo `LEEME.md` para entender los requisitos
3. Ejecuta los tests con `go test` para ver qué funciones necesitas implementar
4. Implementa las funciones en el archivo correspondiente hasta que todos los tests pasen
5. Los tests están diseñados para guiarte en la implementación paso a paso

## Ejercicios

### BÁSICO

#### 1. CALCULATOR - Calculadora Básica
- **Archivo**: `BASIC/01_CALCULATOR/`
- **Dificultad**: ⭐
- **Conceptos**: Funciones básicas, operaciones matemáticas, manejo de errores

#### 2. STRING_UTILS - Utilidades de Strings
- **Archivo**: `BASIC/02_STRING_UTILS/`
- **Dificultad**: ⭐
- **Conceptos**: Manipulación de strings, slices, loops

#### 3. ARRAY_OPERATIONS - Operaciones con Arrays
- **Archivo**: `BASIC/03_ARRAY_OPERATIONS/`
- **Dificultad**: ⭐⭐
- **Conceptos**: Slices, arrays, algoritmos básicos

### INTERMEDIO

#### 4. BANK_ACCOUNT - Sistema de Cuenta Bancaria
- **Archivo**: `INTERMEDIATE/04_BANK_ACCOUNT/`
- **Dificultad**: ⭐⭐
- **Conceptos**: Structs, métodos, encapsulación

#### 5. TODO_LIST - Lista de Tareas
- **Archivo**: `INTERMEDIATE/05_TODO_LIST/`
- **Dificultad**: ⭐⭐
- **Conceptos**: Structs, slices, CRUD operations

#### 6. FILE_PROCESSOR - Procesador de Archivos
- **Archivo**: `INTERMEDIATE/06_FILE_PROCESSOR/`
- **Dificultad**: ⭐⭐⭐
- **Conceptos**: File I/O, error handling, string processing

#### 7. HTTP_CLIENT - Cliente HTTP
- **Archivo**: `INTERMEDIATE/07_HTTP_CLIENT/`
- **Dificultad**: ⭐⭐⭐
- **Conceptos**: HTTP requests, JSON, error handling

### AVANZADO

#### 8. CONCURRENT_DOWNLOADER - Descargador Concurrente
- **Archivo**: `ADVANCED/08_CONCURRENT_DOWNLOADER/`
- **Dificultad**: ⭐⭐⭐⭐
- **Conceptos**: Goroutines, channels, concurrency

#### 9. CACHE_SYSTEM - Sistema de Cache
- **Archivo**: `ADVANCED/09_CACHE_SYSTEM/`
- **Dificultad**: ⭐⭐⭐⭐
- **Conceptos**: Maps, mutex, concurrency, TTL

#### 10. WEB_SCRAPER - Web Scraper
- **Archivo**: `ADVANCED/10_WEB_SCRAPER/`
- **Dificultad**: ⭐⭐⭐⭐⭐
- **Conceptos**: HTTP, HTML parsing, goroutines, channels

### INTELIGENCIA ARTIFICIAL

#### 11. SENTIMENT_ANALYZER - Analizador de Sentimientos
- **Archivo**: `IA/11_SENTIMENT_ANALYZER/`
- **Dificultad**: ⭐⭐
- **Conceptos**: Procesamiento de texto, análisis de sentimientos, diccionarios

#### 12. IMAGE_CLASSIFIER - Clasificador de Imágenes
- **Archivo**: `IA/12_IMAGE_CLASSIFIER/`
- **Dificultad**: ⭐⭐⭐⭐
- **Conceptos**: Visión por computadora, modelos pre-entrenados, GoCV

#### 13. CHATBOT_ENGINE - Motor de Chatbot
- **Archivo**: `IA/13_CHATBOT_ENGINE/`
- **Dificultad**: ⭐⭐⭐⭐⭐
- **Conceptos**: Procesamiento de lenguaje natural, contexto de conversación, APIs externas

## Comandos Útiles

```bash
# Ejecutar tests de un ejercicio específico
cd BASIC/01_CALCULATOR && go test

# Ejecutar tests con verbose output
go test -v

# Ejecutar todos los tests del proyecto
go test ./...

# Verificar coverage
go test -cover
```

## Tips para el Desarrollo TDD

1. **Red**: Ejecuta los tests y ve que fallan
2. **Green**: Escribe el código mínimo para hacer pasar los tests
3. **Refactor**: Mejora el código manteniendo los tests verdes

¡Buena suerte con los ejercicios!