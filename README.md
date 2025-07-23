# Go Practice Exercises - TDD Style

This repository contains Go practice exercises organized by difficulty level. Each exercise has pre-written test cases that you need to pass by implementing the corresponding functions.

## Personal advice

This is a repository for practicing and learning GO; the idea is that you rely on official documentation, forums, Q&A platforms such as Stack Overflow, etc. Try not to let an AI agent solve everything for you. You can rely on it a little, but the idea is that you fully understand what you are developing.

## Project structure

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

## Instructions

1. For each exercise, browse to the corresponding directory
2. Read the `README.md` file to understand the requirements
3. Run the tests with `go test` to see which functions you need to implement
4. Implement the functions in the corresponding file until all tests pass
5. The tests are designed to guide you step by step in the implementation

## Exercises

### BASIC

#### 1. CALCULATOR - Basic calculator
- **Archivo**: `BASIC/01_CALCULATOR/`
- **Dificultad**: ⭐
- **Conceptos**: Basic functions, math operations, error handling

#### 2. STRING_UTILS - Strings utilities
- **Archivo**: `BASIC/02_STRING_UTILS/`
- **Dificultad**: ⭐
- **Conceptos**: Strings handling, slices, loops

#### 3. ARRAY_OPERATIONS - Operations with Arrays
- **Archivo**: `BASIC/03_ARRAY_OPERATIONS/`
- **Dificultad**: ⭐⭐
- **Conceptos**: Slices, arrays, basic algorithms

### INTERMEDIO

#### 4. BANK_ACCOUNT - Bank account system
- **Archivo**: `INTERMEDIATE/04_BANK_ACCOUNT/`
- **Dificultad**: ⭐⭐
- **Conceptos**: Structs, methods, encapsulation

#### 5. TODO_LIST - Tasks list
- **Archivo**: `INTERMEDIATE/05_TODO_LIST/`
- **Dificultad**: ⭐⭐
- **Conceptos**: Structs, slices, CRUD operations

#### 6. FILE_PROCESSOR - File processor
- **Archivo**: `INTERMEDIATE/06_FILE_PROCESSOR/`
- **Dificultad**: ⭐⭐⭐
- **Conceptos**: File I/O, error handling, string processing

#### 7. HTTP_CLIENT - HTTP Client
- **Archivo**: `INTERMEDIATE/07_HTTP_CLIENT/`
- **Dificultad**: ⭐⭐⭐
- **Conceptos**: HTTP requests, JSON, error handling

### AVANZADO

#### 8. CONCURRENT_DOWNLOADER - Concurrent downloader
- **Archivo**: `ADVANCED/08_CONCURRENT_DOWNLOADER/`
- **Dificultad**: ⭐⭐⭐⭐
- **Conceptos**: Goroutines, channels, concurrency

#### 9. CACHE_SYSTEM - Cache system
- **Archivo**: `ADVANCED/09_CACHE_SYSTEM/`
- **Dificultad**: ⭐⭐⭐⭐
- **Conceptos**: Maps, mutex, concurrency, TTL

#### 10. WEB_SCRAPER - Web Scraper
- **Archivo**: `ADVANCED/10_WEB_SCRAPER/`
- **Dificultad**: ⭐⭐⭐⭐⭐
- **Conceptos**: HTTP, HTML parsing, goroutines, channels

### INTELIGENCIA ARTIFICIAL

#### 11. SENTIMENT_ANALYZER - Feeling Analyzer
- **Archivo**: `IA/11_SENTIMENT_ANALYZER/`
- **Dificultad**: ⭐⭐
- **Conceptos**: Text procesing, sentiment analisis, dictionaries

#### 12. IMAGE_CLASSIFIER - Image classifier
- **Archivo**: `IA/12_IMAGE_CLASSIFIER/`
- **Dificultad**: ⭐⭐⭐⭐
- **Conceptos**: Computer vision, pre-trained models, GoCV

#### 13. CHATBOT_ENGINE - Chatbot engine
- **Archivo**: `IA/13_CHATBOT_ENGINE/`
- **Dificultad**: ⭐⭐⭐⭐⭐
- **Conceptos**: Natural Lnguage Process, conversation context, External APIs

## Comandos Útiles

```bash
# Execute tests of specific exercise
cd BASIC/01_CALCULATOR && go test

# Executes tests with verbose output
go test -v

# Execute all project tests
go test ./...

# Check coverage
go test -cover
```

## Tips for TDD

1. **Red**: Run the test and see the fails
2. **Green**: Write the minimal code to pass the test
3. **Refactor**: Refactor the code to improve readability and maintainability

¡Happy coding!
