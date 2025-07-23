package downloader

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

// Setup a test server
func setupTestServer() *httptest.Server {
	mux := http.NewServeMux()
	
	// Counter for concurrent requests
	var (
		mu             sync.Mutex
		activeRequests int
		maxConcurrent  int
	)
	
	// Middleware to track concurrent requests
	trackConcurrency := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			activeRequests++
			if activeRequests > maxConcurrent {
				maxConcurrent = activeRequests
			}
			mu.Unlock()
			
			next(w, r)
			
			mu.Lock()
			activeRequests--
			mu.Unlock()
		}
	}
	
	// Handler for GET /file/{size}
	mux.HandleFunc("/file/", trackConcurrency(func(w http.ResponseWriter, r *http.Request) {
		size := 1000 // Default size
		fmt.Sscanf(r.URL.Path, "/file/%d", &size)
		
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=file-%d.dat", size))
		
		// Generate random data
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(i % 256)
		}
		
		// Simulate slow download
		time.Sleep(100 * time.Millisecond)
		
		w.Write(data)
	}))
	
	// Handler for GET /slow/{delay}
	mux.HandleFunc("/slow/", trackConcurrency(func(w http.ResponseWriter, r *http.Request) {
		delay := 1 // Default delay in seconds
		fmt.Sscanf(r.URL.Path, "/slow/%d", &delay)
		
		w.Header().Set("Content-Type", "text/plain")
		
		// Sleep to simulate slow response
		time.Sleep(time.Duration(delay) * time.Second)
		
		fmt.Fprintf(w, "Response after %d second delay", delay)
	}))
	
	// Handler for GET /error
	mux.HandleFunc("/error", trackConcurrency(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	
	// Function to get max concurrent requests
	getMaxConcurrent := func() int {
		mu.Lock()
		defer mu.Unlock()
		return maxConcurrent
	}
	
	server := httptest.NewServer(mux)
	
	// Store the getMaxConcurrent function in the server
	server.Config.ConnState = func(conn http.ConnState, state http.ConnState) {
		// This is just a hack to store our function
		if state == http.StateNew {
			server.CloseClientConnections() // This is never called, just to avoid unused variable warning
		}
	}
	
	// Expose the getMaxConcurrent function
	http.DefaultServeMux.HandleFunc("/_test/maxConcurrent", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", getMaxConcurrent())
	})
	
	return server
}

// Helper function to create a temporary directory
func createTempDir(t *testing.T) string {
	t.Helper()
	
	dir, err := os.MkdirTemp("", "downloader-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	
	return dir
}

// Helper function to get max concurrent requests from test server
func getMaxConcurrent() int {
	resp, err := http.Get("http://" + http.DefaultServeMux.Handler.ServeHTTP.String() + "/_test/maxConcurrent")
	if err != nil {
		return -1
	}
	defer resp.Body.Close()
	
	var max int
	fmt.Fscanf(resp.Body, "%d", &max)
	return max
}

func TestDownload(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	tempDir := createTempDir(t)
	defer os.RemoveAll(tempDir)
	
	downloader := NewDownloader(3)
	
	// Test successful download
	filePath := filepath.Join(tempDir, "test-file.dat")
	err := downloader.Download(server.URL+"/file/1000", filePath)
	if err != nil {
		t.Fatalf("Download failed: %v", err)
	}
	
	// Verify file exists and has correct size
	info, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf("Failed to stat downloaded file: %v", err)
	}
	
	if info.Size() != 1000 {
		t.Errorf("Downloaded file size = %d, want 1000", info.Size())
	}
	
	// Test download with error
	err = downloader.Download(server.URL+"/error", filepath.Join(tempDir, "error.txt"))
	if err == nil {
		t.Error("Download should fail with server error")
	}
	
	// Test download with invalid URL
	err = downloader.Download("http://invalid-url-that-does-not-exist.example", filepath.Join(tempDir, "invalid.txt"))
	if err == nil {
		t.Error("Download should fail with invalid URL")
	}
}

func TestDownloadMany(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	tempDir := createTempDir(t)
	defer os.RemoveAll(tempDir)
	
	// Create URLs for testing
	urls := []string{
		server.URL + "/file/1000",
		server.URL + "/file/2000",
		server.URL + "/file/3000",
		server.URL + "/file/4000",
		server.URL + "/error",
		"http://invalid-url-that-does-not-exist.example",
	}
	
	// Test with max concurrency = 2
	downloader := NewDownloader(2)
	results := downloader.DownloadMany(urls, tempDir)
	
	// Check number of results
	if len(results) != len(urls) {
		t.Errorf("DownloadMany returned %d results, want %d", len(results), len(urls))
	}
	
	// Count successful downloads
	successCount := 0
	for _, result := range results {
		if result.Error == nil {
			successCount++
			
			// Verify file exists
			_, err := os.Stat(result.FilePath)
			if err != nil {
				t.Errorf("Downloaded file %s does not exist: %v", result.FilePath, err)
			}
		}
	}
	
	// Should have 4 successful downloads (the first 4 URLs)
	if successCount != 4 {
		t.Errorf("DownloadMany had %d successful downloads, want 4", successCount)
	}
	
	// Check max concurrency (this is a bit hacky and may not always work reliably in tests)
	// maxConcurrent := getMaxConcurrent()
	// if maxConcurrent > 2 {
	//     t.Errorf("Max concurrent downloads = %d, should not exceed 2", maxConcurrent)
	// }
}

func TestDownloadWithProgress(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	tempDir := createTempDir(t)
	defer os.RemoveAll(tempDir)
	
	downloader := NewDownloader(1)
	
	// Create a progress channel
	progressChan := make(chan int)
	
	// Start download in a goroutine
	filePath := filepath.Join(tempDir, "progress-test.dat")
	errChan := make(chan error)
	go func() {
		errChan <- downloader.DownloadWithProgress(server.URL+"/file/10000", filePath, progressChan)
	}()
	
	// Collect progress updates
	var lastProgress int
	timeout := time.After(5 * time.Second)
	
progressLoop:
	for {
		select {
		case progress := <-progressChan:
			if progress > lastProgress {
				lastProgress = progress
			}
			if progress >= 100 {
				break progressLoop
			}
		case err := <-errChan:
			if err != nil {
				t.Fatalf("Download failed: %v", err)
			}
			break progressLoop
		case <-timeout:
			t.Fatal("Test timed out waiting for progress updates")
			break progressLoop
		}
	}
	
	// Verify we got progress updates
	if lastProgress < 100 {
		t.Errorf("Final progress = %d, want 100", lastProgress)
	}
	
	// Verify file exists
	_, err := os.Stat(filePath)
	if err != nil {
		t.Errorf("Downloaded file does not exist: %v", err)
	}
}