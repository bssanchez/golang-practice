package fileprocessor

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// ReadFile reads a file and returns its lines as a slice of strings
func ReadFile(filename string) ([]string, error) {
	// TODO: Implement this function
	return nil, nil
}

// WriteFile writes a slice of strings to a file, one line per element
func WriteFile(filename string, lines []string) error {
	// TODO: Implement this function
	return nil
}

// CountWords counts the number of words in a file
func CountWords(filename string) (int, error) {
	// TODO: Implement this function
	return 0, nil
}

// FindPattern finds lines in a file that match a regex pattern
func FindPattern(filename, pattern string) ([]string, error) {
	// TODO: Implement this function
	return nil, nil
}

// ReplaceInFile replaces all occurrences of 'old' with 'new' in a file
// Returns the number of replacements made
func ReplaceInFile(filename, old, new string) (int, error) {
	// TODO: Implement this function
	return 0, nil
}