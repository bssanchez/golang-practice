# Exercise 10: Web Scraper

## Description
Implement a concurrent web scraper that can extract information from web pages and follow links.

## Requirements
1. Implement the following structures and functions in the `scraper.go` file:
   - Struct `ScrapedPage` with fields for URL, title, links, and extracted text
   - Struct `Scraper` with fields for maximum depth, maximum number of pages, and allowed domain
   - Method `NewScraper(maxDepth, maxPages int, allowedDomain string) *Scraper` - Constructor
   - Method `Scrape(startURL string) ([]ScrapedPage, error)` - Starts scraping from a URL
   - Method `ExtractLinks(html string, baseURL string) ([]string, error)` - Extracts links from HTML
   - Method `ExtractTitle(html string) string` - Extracts the title from an HTML page
   - Method `ExtractText(html string) string` - Extracts the main text from an HTML page

2. Considerations:
   - Use goroutines for concurrent scraping
   - Implement a mechanism to limit concurrency
   - Respect robots.txt and add delays between requests
   - Handle relative and absolute URLs correctly
   - Implement a mechanism to avoid cycles (don't visit the same URL twice)
   - Limit scraping to the allowed domain

## Tests
Run `go test` to verify your implementation.

## Note
The tests use a mock HTTP server to avoid external dependencies.