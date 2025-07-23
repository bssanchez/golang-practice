# Exercise 2: String Utilities

## Description
Implement a series of functions to manipulate strings in Go.

## Requirements
1. Implement the following functions in the `stringutils.go` file:
   - `Reverse(s string) string` - Reverses a string
   - `IsPalindrome(s string) bool` - Checks if a string is a palindrome
   - `CountOccurrences(s, substr string) int` - Counts occurrences of a substring
   - `ToTitleCase(s string) string` - Converts the first letter of each word to uppercase
   - `RemoveDuplicateChars(s string) string` - Removes duplicate characters

2. Considerations:
   - Functions should handle empty strings correctly
   - For `IsPalindrome`, ignore case and spaces
   - For `ToTitleCase`, consider that words are separated by spaces

## Tests
Run `go test` to verify your implementation.