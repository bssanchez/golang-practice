package sentiment

import (
	"path/filepath"
	"testing"
)

func TestLoadSentimentDictionary(t *testing.T) {
	// Paths to test files
	positiveFile := filepath.Join("data", "positive_words.txt")
	negativeFile := filepath.Join("data", "negative_words.txt")
	
	// Load dictionary
	dict, err := LoadSentimentDictionary(positiveFile, negativeFile)
	if err != nil {
		t.Fatalf("Failed to load sentiment dictionary: %v", err)
	}
	
	// Check that dictionary is not nil
	if dict == nil {
		t.Fatal("LoadSentimentDictionary returned nil dictionary")
	}
	
	// Check that positive words were loaded
	if len(dict.PositiveWords) == 0 {
		t.Error("No positive words were loaded")
	}
	
	// Check that negative words were loaded
	if len(dict.NegativeWords) == 0 {
		t.Error("No negative words were loaded")
	}
	
	// Check specific words
	if !dict.PositiveWords["bueno"] {
		t.Error("Expected 'bueno' to be in positive words")
	}
	
	if !dict.NegativeWords["malo"] {
		t.Error("Expected 'malo' to be in negative words")
	}
	
	// Test with non-existent files
	_, err = LoadSentimentDictionary("nonexistent.txt", negativeFile)
	if err == nil {
		t.Error("LoadSentimentDictionary should return error for non-existent positive file")
	}
	
	_, err = LoadSentimentDictionary(positiveFile, "nonexistent.txt")
	if err == nil {
		t.Error("LoadSentimentDictionary should return error for non-existent negative file")
	}
}

func TestGetSentimentScore(t *testing.T) {
	// Create a test dictionary
	dict := &SentimentDictionary{
		PositiveWords: map[string]bool{
			"bueno":      true,
			"excelente":  true,
			"fantástico": true,
		},
		NegativeWords: map[string]bool{
			"malo":     true,
			"terrible": true,
			"horrible": true,
		},
	}
	
	tests := []struct {
		name     string
		text     string
		expected float64
	}{
		{"positive text", "Este producto es bueno y fantástico", 2.0},
		{"negative text", "Este servicio es malo y terrible", -2.0},
		{"mixed text", "El producto es bueno pero el servicio es malo", 0.0},
		{"neutral text", "Este es un producto", 0.0},
		{"text with negation", "Este producto no es malo", 1.0}, // Negation should flip sentiment
		{"empty text", "", 0.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := GetSentimentScore(tt.text, dict)
			if score != tt.expected {
				t.Errorf("GetSentimentScore(%q) = %f, want %f", tt.text, score, tt.expected)
			}
		})
	}
}

func TestAnalyzeSentiment(t *testing.T) {
	// Create a test dictionary
	dict := &SentimentDictionary{
		PositiveWords: map[string]bool{
			"bueno":      true,
			"excelente":  true,
			"fantástico": true,
		},
		NegativeWords: map[string]bool{
			"malo":     true,
			"terrible": true,
			"horrible": true,
		},
	}
	
	tests := []struct {
		name         string
		text         string
		expectedSent string
	}{
		{"positive text", "Este producto es bueno y fantástico", "positive"},
		{"negative text", "Este servicio es malo y terrible", "negative"},
		{"neutral text", "Este es un producto", "neutral"},
		{"text with negation", "Este producto no es malo", "positive"},
		{"empty text", "", "neutral"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AnalyzeSentiment(tt.text, dict)
			if result.Sentiment != tt.expectedSent {
				t.Errorf("AnalyzeSentiment(%q) = %s, want %s", tt.text, result.Sentiment, tt.expectedSent)
			}
			
			// Check that the original text is preserved
			if result.Text != tt.text {
				t.Errorf("AnalyzeSentiment did not preserve original text: got %q, want %q", result.Text, tt.text)
			}
			
			// Check that score is consistent with sentiment
			if (result.Sentiment == "positive" && result.Score <= 0) ||
			   (result.Sentiment == "negative" && result.Score >= 0) ||
			   (result.Sentiment == "neutral" && result.Score != 0) {
				t.Errorf("Inconsistent sentiment and score: sentiment=%s, score=%f", result.Sentiment, result.Score)
			}
		})
	}
}

func TestAnalyzeMultipleTexts(t *testing.T) {
	// Create a test dictionary
	dict := &SentimentDictionary{
		PositiveWords: map[string]bool{
			"bueno":      true,
			"excelente":  true,
			"fantástico": true,
		},
		NegativeWords: map[string]bool{
			"malo":     true,
			"terrible": true,
			"horrible": true,
		},
	}
	
	texts := []string{
		"Este producto es bueno",
		"Este servicio es malo",
		"Esto es neutral",
	}
	
	results := AnalyzeMultipleTexts(texts, dict)
	
	// Check number of results
	if len(results) != len(texts) {
		t.Errorf("AnalyzeMultipleTexts returned %d results, want %d", len(results), len(texts))
	}
	
	// Check individual results
	expectedSentiments := []string{"positive", "negative", "neutral"}
	for i, result := range results {
		if result.Text != texts[i] {
			t.Errorf("Result %d: text = %q, want %q", i, result.Text, texts[i])
		}
		
		if result.Sentiment != expectedSentiments[i] {
			t.Errorf("Result %d: sentiment = %s, want %s", i, result.Sentiment, expectedSentiments[i])
		}
	}
}