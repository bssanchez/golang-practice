package stringutils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"simple word", "hello", "olleh"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"with special chars", "a,b$c!", "!c$b,a"},
		{"with numbers", "abc123", "321cba"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single character", "a", true},
		{"simple palindrome", "radar", true},
		{"mixed case palindrome", "Radar", true},
		{"palindrome with spaces", "a man a plan a canal panama", true},
		{"palindrome with punctuation", "A man, a plan, a canal: Panama", true},
		{"non-palindrome", "hello", false},
		{"almost palindrome", "almosttsomla", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		substr   string
		expected int
	}{
		{"empty string", "", "a", 0},
		{"empty substring", "hello", "", 0},
		{"no occurrences", "hello", "z", 0},
		{"single occurrence", "hello", "h", 1},
		{"multiple occurrences", "hello hello", "hello", 2},
		{"overlapping occurrences", "abababa", "aba", 2},
		{"case sensitive", "Hello hello", "hello", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountOccurrences(tt.s, tt.substr)
			if result != tt.expected {
				t.Errorf("CountOccurrences(%q, %q) = %d; want %d", tt.s, tt.substr, result, tt.expected)
			}
		})
	}
}

func TestToTitleCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single word", "hello", "Hello"},
		{"multiple words", "hello world", "Hello World"},
		{"already title case", "Hello World", "Hello World"},
		{"mixed case", "hELLO wORLD", "Hello World"},
		{"with numbers", "hello 123 world", "Hello 123 World"},
		{"with extra spaces", "  hello  world  ", "  Hello  World  "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToTitleCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToTitleCase(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRemoveDuplicateChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"no duplicates", "abcdef", "abcdef"},
		{"with duplicates", "hello", "helo"},
		{"all duplicates", "aaaaaa", "a"},
		{"with spaces", "hello world", "helo wrd"},
		{"case sensitive", "aAbBcC", "aAbBcC"},
		{"with special chars", "a!b!c!", "a!bc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveDuplicateChars(tt.input)
			if result != tt.expected {
				t.Errorf("RemoveDuplicateChars(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}