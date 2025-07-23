# Exercise 6: File Processor

## Description
Implement a file processor that can read, write, and analyze text files.

## Requirements
1. Implement the following functions in the `fileprocessor.go` file:
   - `ReadFile(filename string) ([]string, error)` - Reads a file and returns its lines
   - `WriteFile(filename string, lines []string) error` - Writes lines to a file
   - `CountWords(filename string) (int, error)` - Counts words in a file
   - `FindPattern(filename, pattern string) ([]string, error)` - Finds lines that match a pattern
   - `ReplaceInFile(filename, old, new string) (int, error)` - Replaces text and returns the number of replacements

2. Considerations:
   - Handle file errors correctly (doesn't exist, permissions, etc.)
   - For `FindPattern`, use regular expressions
   - For `CountWords`, consider that words are separated by spaces, tabs, or line breaks

## Tests
Run `go test` to verify your implementation.

## Note
For tests, temporary files will be created in the current directory.