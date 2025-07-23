# Go Practice Exercises - TDD Style

This repository contains Go practice exercises organized by difficulty level. Each exercise has pre-written test cases that you need to pass by implementing the corresponding functions.

## Personal advice

This is a repository for practicing and learning GO; the idea is that you rely on official documentation, forums, Q&A platforms such as Stack Overflow, etc. Try not to let an AI agent solve everything for you. You can rely on it a little, but the idea is that you fully understand what you are developing.

## Project structure

```
./
├── 01_BASIC/
│   ├── 01_CALCULATOR/
│   ├── 02_STRING_UTILS/
│   └── 03_ARRAY_OPERATIONS/
├── 02_INTERMEDIATE/
│   ├── 04_BANK_ACCOUNT/
│   ├── 05_TODO_LIST/
│   ├── 06_FILE_PROCESSOR/
│   └── 07_HTTP_CLIENT/
├── 03_ADVANCED/
│   ├── 08_CONCURRENT_DOWNLOADER/
│   ├── 09_CACHE_SYSTEM/
│   └── 10_WEB_SCRAPER/
└── 04_AI/
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
- **Files**: `01_BASIC/01_CALCULATOR/`
- **Difficulty**: ⭐
- **Concepts**: Basic functions, math operations, error handling

#### 2. STRING_UTILS - Strings utilities
- **Files**: `01_BASIC/02_STRING_UTILS/`
- **Difficulty**: ⭐
- **Concepts**: Strings handling, slices, loops

#### 3. ARRAY_OPERATIONS - Operations with Arrays
- **Files**: `01_BASIC/03_ARRAY_OPERATIONS/`
- **Difficulty**: ⭐⭐
- **Concepts**: Slices, arrays, basic algorithms

### INTERMEDIO

#### 4. BANK_ACCOUNT - Bank account system
- **Files**: `02_INTERMEDIATE/04_BANK_ACCOUNT/`
- **Difficulty**: ⭐⭐
- **Concepts**: Structs, methods, encapsulation

#### 5. TODO_LIST - Tasks list
- **Files**: `02_INTERMEDIATE/05_TODO_LIST/`
- **Difficulty**: ⭐⭐
- **Concepts**: Structs, slices, CRUD operations

#### 6. FILE_PROCESSOR - File processor
- **Files**: `02_INTERMEDIATE/06_FILE_PROCESSOR/`
- **Difficulty**: ⭐⭐⭐
- **Concepts**: File I/O, error handling, string processing

#### 7. HTTP_CLIENT - HTTP Client
- **Files**: `02_INTERMEDIATE/07_HTTP_CLIENT/`
- **Difficulty**: ⭐⭐⭐
- **Concepts**: HTTP requests, JSON, error handling

### AVANZADO

#### 8. CONCURRENT_DOWNLOADER - Concurrent downloader
- **Files**: `03_ADVANCED/08_CONCURRENT_DOWNLOADER/`
- **Difficulty**: ⭐⭐⭐⭐
- **Concepts**: Goroutines, channels, concurrency

#### 9. CACHE_SYSTEM - Cache system
- **Files**: `03_ADVANCED/09_CACHE_SYSTEM/`
- **Difficulty**: ⭐⭐⭐⭐
- **Concepts**: Maps, mutex, concurrency, TTL

#### 10. WEB_SCRAPER - Web Scraper
- **Files**: `03_ADVANCED/10_WEB_SCRAPER/`
- **Difficulty**: ⭐⭐⭐⭐⭐
- **Concepts**: HTTP, HTML parsing, goroutines, channels

### INTELIGENCIA ARTIFICIAL

#### 11. SENTIMENT_ANALYZER - Feeling Analyzer
- **Files**: `04_AI/11_SENTIMENT_ANALYZER/`
- **Difficulty**: ⭐⭐
- **Concepts**: Text procesing, sentiment analisis, dictionaries

#### 12. IMAGE_CLASSIFIER - Image classifier
- **Files**: `04_AI/12_IMAGE_CLASSIFIER/`
- **Difficulty**: ⭐⭐⭐⭐
- **Concepts**: Computer vision, pre-trained models, GoCV

#### 13. CHATBOT_ENGINE - Chatbot engine
- **Files**: `04_AI/13_CHATBOT_ENGINE/`
- **Difficulty**: ⭐⭐⭐⭐⭐
- **Concepts**: Natural Lnguage Process, conversation context, External APIs

## Comandos Útiles

```bash
# Execute tests of specific exercise
cd 01_BASIC/01_CALCULATOR && go test

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
