package scraper

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ScrapedPage contains information extracted from a web page
type ScrapedPage struct {
	URL     string   // URL of the page
	Title   string   // Title of the page
	Text    string   // Main text of the page
	Links   []string // Links found on the page
	Depth   int      // Depth of the page in the scraping tree
}

// Scraper implements a concurrent web scraper
type Scraper struct {
	MaxDepth      int    // Maximum scraping depth
	MaxPages      int    // Maximum number of pages to scrape
	AllowedDomain string // Allowed domain for scraping
}

// NewScraper creates a new scraper with the specified limits
func NewScraper(maxDepth, maxPages int, allowedDomain string) *Scraper {
	// TODO: Implement this function
	return nil
}

// Scrape starts the scraping process from an initial URL
// Returns the scraped pages and a possible error
func (s *Scraper) Scrape(startURL string) ([]ScrapedPage, error) {
	// TODO: Implement this function
	return nil, nil
}

// ExtractLinks extracts all links from an HTML document
// Converts relative URLs to absolute using the base URL
func (s *Scraper) ExtractLinks(html string, baseURL string) ([]string, error) {
	// TODO: Implement this function
	return nil, nil
}

// ExtractTitle extracts the title from an HTML document
func (s *Scraper) ExtractTitle(html string) string {
	// TODO: Implement this function
	return ""
}

// ExtractText extracts the main text from an HTML document
// Removes HTML tags, scripts, styles, etc.
func (s *Scraper) ExtractText(html string) string {
	// TODO: Implement this function
	return ""
}