package fileprocessor

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// Helper function to create a temporary file with content
func createTempFile(t *testing.T, content string) string {
	t.Helper()
	
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	
	// Write content to the file
	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	
	// Close the file
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}
	
	return tmpfile.Name()
}

// Helper function to clean up temporary files
func cleanupTempFile(t *testing.T, filename string) {
	t.Helper()
	os.Remove(filename)
}

func TestReadFile(t *testing.T) {
	// Create a temporary file
	content := "Line 1\nLine 2\nLine 3"
	filename := createTempFile(t, content)
	defer cleanupTempFile(t, filename)
	
	// Test reading the file
	lines, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile(%q) returned error: %v", filename, err)
	}
	
	expected := []string{"Line 1", "Line 2", "Line 3"}
	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("ReadFile(%q) = %v, want %v", filename, lines, expected)
	}
	
	// Test reading non-existent file
	_, err = ReadFile("nonexistent.txt")
	if err == nil {
		t.Error("ReadFile on non-existent file should return error")
	}
}

func TestWriteFile(t *testing.T) {
	// Create a temporary filename
	tmpdir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpdir)
	
	filename := filepath.Join(tmpdir, "output.txt")
	
	// Test writing to the file
	lines := []string{"Hello", "World", "Go"}
	err = WriteFile(filename, lines)
	if err != nil {
		t.Fatalf("WriteFile(%q) returned error: %v", filename, err)
	}
	
	// Read back the file to verify
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}
	
	expected := "Hello\nWorld\nGo"
	if strings.TrimSpace(string(content)) != expected {
		t.Errorf("WriteFile(%q) wrote %q, want %q", filename, string(content), expected)
	}
	
	// Test writing to a directory (should fail)
	err = WriteFile(tmpdir, lines)
	if err == nil {
		t.Error("WriteFile to a directory should return error")
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected int
	}{
		{"empty file", "", 0},
		{"single word", "hello", 1},
		{"multiple words", "hello world go", 3},
		{"multiple lines", "hello world\ngo programming\nis fun", 5},
		{"with punctuation", "hello, world! This is a test.", 6},
		{"with extra spaces", "  hello  world  ", 2},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := createTempFile(t, tt.content)
			defer cleanupTempFile(t, filename)
			
			count, err := CountWords(filename)
			if err != nil {
				t.Fatalf("CountWords(%q) returned error: %v", filename, err)
			}
			
			if count != tt.expected {
				t.Errorf("CountWords(%q) = %d, want %d", filename, count, tt.expected)
			}
		})
	}
	
	// Test counting words in non-existent file
	_, err := CountWords("nonexistent.txt")
	if err == nil {
		t.Error("CountWords on non-existent file should return error")
	}
}

func TestFindPattern(t *testing.T) {
	// Create a temporary file
	content := `Line 1 with apple
Line 2 with banana
Line 3 with apple and banana
Line 4 with orange
Line 5 with APPLE in caps`
	
	filename := createTempFile(t, content)
	defer cleanupTempFile(t, filename)
	
	// Test finding a pattern
	lines, err := FindPattern(filename, "apple")
	if err != nil {
		t.Fatalf("FindPattern(%q, %q) returned error: %v", filename, "apple", err)
	}
	
	expected := []string{"Line 1 with apple", "Line 3 with apple and banana"}
	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("FindPattern(%q, %q) = %v, want %v", filename, "apple", lines, expected)
	}
	
	// Test case-insensitive pattern
	lines, err = FindPattern(filename, "(?i)apple")
	if err != nil {
		t.Fatalf("FindPattern(%q, %q) returned error: %v", filename, "(?i)apple", err)
	}
	
	expected = []string{"Line 1 with apple", "Line 3 with apple and banana", "Line 5 with APPLE in caps"}
	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("FindPattern(%q, %q) = %v, want %v", filename, "(?i)apple", lines, expected)
	}
	
	// Test pattern with no matches
	lines, err = FindPattern(filename, "pear")
	if err != nil {
		t.Fatalf("FindPattern(%q, %q) returned error: %v", filename, "pear", err)
	}
	
	if len(lines) != 0 {
		t.Errorf("FindPattern(%q, %q) should return empty slice, got %v", filename, "pear", lines)
	}
	
	// Test invalid pattern
	_, err = FindPattern(filename, "[")
	if err == nil {
		t.Error("FindPattern with invalid pattern should return error")
	}
	
	// Test non-existent file
	_, err = FindPattern("nonexistent.txt", "apple")
	if err == nil {
		t.Error("FindPattern on non-existent file should return error")
	}
}

func TestReplaceInFile(t *testing.T) {
	// Create a temporary file
	content := "apple orange apple banana apple"
	filename := createTempFile(t, content)
	defer cleanupTempFile(t, filename)
	
	// Test replacing text
	count, err := ReplaceInFile(filename, "apple", "pear")
	if err != nil {
		t.Fatalf("ReplaceInFile(%q, %q, %q) returned error: %v", filename, "apple", "pear", err)
	}
	
	if count != 3 {
		t.Errorf("ReplaceInFile(%q, %q, %q) = %d, want 3", filename, "apple", "pear", count)
	}
	
	// Read back the file to verify
	lines, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read modified file: %v", err)
	}
	
	expected := "pear orange pear banana pear"
	if len(lines) != 1 || lines[0] != expected {
		t.Errorf("After ReplaceInFile, file content = %q, want %q", lines, expected)
	}
	
	// Test replacing with no matches
	count, err = ReplaceInFile(filename, "apple", "grape")
	if err != nil {
		t.Fatalf("ReplaceInFile(%q, %q, %q) returned error: %v", filename, "apple", "grape", err)
	}
	
	if count != 0 {
		t.Errorf("ReplaceInFile with no matches should return 0, got %d", count)
	}
	
	// Test non-existent file
	_, err = ReplaceInFile("nonexistent.txt", "apple", "pear")
	if err == nil {
		t.Error("ReplaceInFile on non-existent file should return error")
	}
}