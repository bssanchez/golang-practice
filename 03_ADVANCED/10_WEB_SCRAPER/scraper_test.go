package scraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// Setup a test server with mock pages
func setupTestServer() *httptest.Server {
	mux := http.NewServeMux()
	
	// Home page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Home Page</title>
			</head>
			<body>
				<h1>Welcome to the Test Site</h1>
				<p>This is a test site for the web scraper.</p>
				<ul>
					<li><a href="/page1">Page 1</a></li>
					<li><a href="/page2">Page 2</a></li>
					<li><a href="/page3">Page 3</a></li>
					<li><a href="https://external.example.com">External Link</a></li>
				</ul>
			</body>
			</html>
		`)
	})
	
	// Page 1
	mux.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Page 1</title>
			</head>
			<body>
				<h1>Page 1</h1>
				<p>This is page 1 of the test site.</p>
				<ul>
					<li><a href="/">Home</a></li>
					<li><a href="/page2">Page 2</a></li>
					<li><a href="/subpage1">Subpage 1</a></li>
				</ul>
			</body>
			</html>
		`)
	})
	
	// Page 2
	mux.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Page 2</title>
			</head>
			<body>
				<h1>Page 2</h1>
				<p>This is page 2 of the test site.</p>
				<ul>
					<li><a href="/">Home</a></li>
					<li><a href="/page1">Page 1</a></li>
					<li><a href="/page3">Page 3</a></li>
				</ul>
			</body>
			</html>
		`)
	})
	
	// Page 3
	mux.HandleFunc("/page3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Page 3</title>
			</head>
			<body>
				<h1>Page 3</h1>
				<p>This is page 3 of the test site.</p>
				<ul>
					<li><a href="/">Home</a></li>
					<li><a href="/page2">Page 2</a></li>
					<li><a href="/subpage2">Subpage 2</a></li>
				</ul>
			</body>
			</html>
		`)
	})
	
	// Subpage 1
	mux.HandleFunc("/subpage1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Subpage 1</title>
			</head>
			<body>
				<h1>Subpage 1</h1>
				<p>This is a subpage of page 1.</p>
				<ul>
					<li><a href="/page1">Back to Page 1</a></li>
					<li><a href="/">Home</a></li>
				</ul>
			</body>
			</html>
		`)
	})
	
	// Subpage 2
	mux.HandleFunc("/subpage2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Subpage 2</title>
			</head>
			<body>
				<h1>Subpage 2</h1>
				<p>This is a subpage of page 3.</p>
				<ul>
					<li><a href="/page3">Back to Page 3</a></li>
					<li><a href="/">Home</a></li>
				</ul>
			</body>
			</html>
		`)
	})
	
	// Robots.txt
	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			User-agent: *
			Disallow: /private/
		`)
	})
	
	// Private page (should be blocked by robots.txt)
	mux.HandleFunc("/private/secret", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Secret Page</title>
			</head>
			<body>
				<h1>Secret Page</h1>
				<p>This page should not be scraped due to robots.txt.</p>
			</body>
			</html>
		`)
	})
	
	return httptest.NewServer(mux)
}

func TestExtractLinks(t *testing.T) {
	scraper := NewScraper(3, 10, "example.com")
	
	html := `
		<html>
		<body>
			<a href="https://example.com/page1">Page 1</a>
			<a href="/page2">Page 2</a>
			<a href="page3">Page 3</a>
			<a href="https://other.com/page4">External Page</a>
		</body>
		</html>
	`
	
	baseURL := "https://example.com"
	links, err := scraper.ExtractLinks(html, baseURL)
	
	if err != nil {
		t.Fatalf("ExtractLinks returned error: %v", err)
	}
	
	expectedLinks := []string{
		"https://example.com/page1",
		"https://example.com/page2",
		"https://example.com/page3",
		"https://other.com/page4",
	}
	
	if len(links) != len(expectedLinks) {
		t.Errorf("ExtractLinks returned %d links, want %d", len(links), len(expectedLinks))
	}
	
	// Check that all expected links are present
	for _, expected := range expectedLinks {
		found := false
		for _, link := range links {
			if link == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("ExtractLinks did not return expected link: %s", expected)
		}
	}
}

func TestExtractTitle(t *testing.T) {
	scraper := NewScraper(3, 10, "example.com")
	
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			"simple title",
			`<html><head><title>Test Title</title></head><body></body></html>`,
			"Test Title",
		},
		{
			"no title",
			`<html><head></head><body></body></html>`,
			"",
		},
		{
			"empty title",
			`<html><head><title></title></head><body></body></html>`,
			"",
		},
		{
			"title with HTML",
			`<html><head><title>Test <b>Bold</b> Title</title></head><body></body></html>`,
			"Test Bold Title",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			title := scraper.ExtractTitle(tt.html)
			if title != tt.expected {
				t.Errorf("ExtractTitle(%q) = %q, want %q", tt.html, title, tt.expected)
			}
		})
	}
}

func TestExtractText(t *testing.T) {
	scraper := NewScraper(3, 10, "example.com")
	
	html := `
		<html>
		<head><title>Test Page</title></head>
		<body>
			<h1>Welcome to the Test Page</h1>
			<p>This is a paragraph of text.</p>
			<div>
				<p>This is another paragraph.</p>
				<script>alert('This should not be included');</script>
				<style>.hidden { display: none; }</style>
			</div>
		</body>
		</html>
	`
	
	text := scraper.ExtractText(html)
	
	// Check that the text contains the main content
	if !strings.Contains(text, "Welcome to the Test Page") {
		t.Error("ExtractText should include the h1 content")
	}
	
	if !strings.Contains(text, "This is a paragraph of text") {
		t.Error("ExtractText should include paragraph content")
	}
	
	if !strings.Contains(text, "This is another paragraph") {
		t.Error("ExtractText should include nested paragraph content")
	}
	
	// Check that script and style content is excluded
	if strings.Contains(text, "alert") {
		t.Error("ExtractText should not include script content")
	}
	
	if strings.Contains(text, "display: none") {
		t.Error("ExtractText should not include style content")
	}
}

func TestScrape(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	// Parse the server URL to get the host
	serverURL, _ := url.Parse(server.URL)
	
	// Create a scraper limited to the test server's domain
	scraper := NewScraper(2, 5, serverURL.Host)
	
	// Scrape the test site
	pages, err := scraper.Scrape(server.URL)
	if err != nil {
		t.Fatalf("Scrape returned error: %v", err)
	}
	
	// Should have scraped 5 pages (maxPages = 5)
	if len(pages) != 5 {
		t.Errorf("Scrape returned %d pages, want 5", len(pages))
	}
	
	// Check that we have the home page
	foundHome := false
	for _, page := range pages {
		if page.URL == server.URL+"/" || page.URL == server.URL {
			foundHome = true
			if page.Title != "Home Page" {
				t.Errorf("Home page title = %q, want %q", page.Title, "Home Page")
			}
			break
		}
	}
	
	if !foundHome {
		t.Error("Scrape did not return the home page")
	}
	
	// Check that we don't have any pages from other domains
	for _, page := range pages {
		pageURL, _ := url.Parse(page.URL)
		if pageURL.Host != serverURL.Host {
			t.Errorf("Scrape returned page from external domain: %s", page.URL)
		}
	}
	
	// Check that we don't have any private pages (blocked by robots.txt)
	for _, page := range pages {
		if strings.Contains(page.URL, "/private/") {
			t.Errorf("Scrape returned page that should be blocked by robots.txt: %s", page.URL)
		}
	}
}

func TestScrapeWithDepthLimit(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	// Parse the server URL to get the host
	serverURL, _ := url.Parse(server.URL)
	
	// Create a scraper with depth limit 1 (only the home page and direct links)
	scraper := NewScraper(1, 10, serverURL.Host)
	
	// Scrape the test site
	pages, err := scraper.Scrape(server.URL)
	if err != nil {
		t.Fatalf("Scrape returned error: %v", err)
	}
	
	// Should have scraped 4 pages (home + 3 direct links)
	if len(pages) != 4 {
		t.Errorf("Scrape with depth 1 returned %d pages, want 4", len(pages))
	}
	
	// Check that we don't have any subpages (which are at depth 2)
	for _, page := range pages {
		if strings.Contains(page.URL, "/subpage") {
			t.Errorf("Scrape with depth 1 returned page at depth 2: %s", page.URL)
		}
	}
}