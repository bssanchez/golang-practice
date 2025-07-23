package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// Test data structure
type TestData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Setup a test server
func setupTestServer() *httptest.Server {
	mux := http.NewServeMux()
	
	// Handler for GET /hello
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprint(w, "Hello, World!")
	})
	
	// Handler for GET /json
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TestData{Name: "test", Value: 42})
	})
	
	// Handler for POST /echo
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		// Read body and echo it back
		defer r.Body.Close()
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		fmt.Fprintf(w, "Received: %s", body)
	})
	
	// Handler for GET /download
	mux.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Disposition", "attachment; filename=test.txt")
		fmt.Fprint(w, "This is a test file for download.")
	})
	
	// Handler for GET /slow
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Sleep for 2 seconds to simulate a slow response
		time.Sleep(2 * time.Second)
		fmt.Fprint(w, "This response was slow")
	})
	
	// Handler for GET /error
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	})
	
	return httptest.NewServer(mux)
}

func TestGet(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	// Test successful GET
	body, err := Get(server.URL + "/hello")
	if err != nil {
		t.Fatalf("Get(%q) returned error: %v", server.URL+"/hello", err)
	}
	
	if string(body) != "Hello, World!" {
		t.Errorf("Get(%q) = %q, want %q", server.URL+"/hello", string(body), "Hello, World!")
	}
	
	// Test GET with error status code
	_, err = Get(server.URL + "/error")
	if err == nil {
		t.Errorf("Get(%q) should return error for status 500", server.URL+"/error")
	}
	
	// Test GET with invalid URL
	_, err = Get("http://invalid-url-that-does-not-exist.example")
	if err == nil {
		t.Error("Get with invalid URL should return error")
	}
}

func TestGetJSON(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	// Test successful JSON GET
	var data TestData
	err := GetJSON(server.URL+"/json", &data)
	if err != nil {
		t.Fatalf("GetJSON(%q) returned error: %v", server.URL+"/json", err)
	}
	
	if data.Name != "test" || data.Value != 42 {
		t.Errorf("GetJSON(%q) = %+v, want {Name:\"test\", Value:42}", server.URL+"/json", data)
	}
	
	// Test JSON GET with error status code
	err = GetJSON(server.URL+"/error", &data)
	if err == nil {
		t.Errorf("GetJSON(%q) should return error for status 500", server.URL+"/error")
	}
	
	// Test JSON GET with invalid URL
	err = GetJSON("http://invalid-url-that-does-not-exist.example", &data)
	if err == nil {
		t.Error("GetJSON with invalid URL should return error")
	}
	
	// Test JSON GET with non-JSON response
	err = GetJSON(server.URL+"/hello", &data)
	if err == nil {
		t.Error("GetJSON with non-JSON response should return error")
	}
}

func TestPost(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	// Test successful POST
	body := []byte("test data")
	response, err := Post(server.URL+"/echo", "text/plain", body)
	if err != nil {
		t.Fatalf("Post(%q) returned error: %v", server.URL+"/echo", err)
	}
	
	expected := "Received: test data"
	if string(response) != expected {
		t.Errorf("Post(%q) = %q, want %q", server.URL+"/echo", string(response), expected)
	}
	
	// Test POST with error status code
	_, err = Post(server.URL+"/error", "text/plain", body)
	if err == nil {
		t.Errorf("Post(%q) should return error for status 500", server.URL+"/error")
	}
	
	// Test POST with invalid URL
	_, err = Post("http://invalid-url-that-does-not-exist.example", "text/plain", body)
	if err == nil {
		t.Error("Post with invalid URL should return error")
	}
}

func TestDownloadFile(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	// Create a temporary directory
	tmpdir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpdir)
	
	// Test successful download
	filepath := filepath.Join(tmpdir, "downloaded.txt")
	err = DownloadFile(server.URL+"/download", filepath)
	if err != nil {
		t.Fatalf("DownloadFile(%q, %q) returned error: %v", server.URL+"/download", filepath, err)
	}
	
	// Verify the file content
	content, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}
	
	expected := "This is a test file for download."
	if string(content) != expected {
		t.Errorf("Downloaded file content = %q, want %q", string(content), expected)
	}
	
	// Test download with error status code
	err = DownloadFile(server.URL+"/error", filepath)
	if err == nil {
		t.Errorf("DownloadFile(%q) should return error for status 500", server.URL+"/error")
	}
	
	// Test download with invalid URL
	err = DownloadFile("http://invalid-url-that-does-not-exist.example", filepath)
	if err == nil {
		t.Error("DownloadFile with invalid URL should return error")
	}
}

func TestFetchWithTimeout(t *testing.T) {
	server := setupTestServer()
	defer server.Close()
	
	// Test successful fetch with sufficient timeout
	body, err := FetchWithTimeout(server.URL+"/hello", 1*time.Second)
	if err != nil {
		t.Fatalf("FetchWithTimeout(%q, 1s) returned error: %v", server.URL+"/hello", err)
	}
	
	if string(body) != "Hello, World!" {
		t.Errorf("FetchWithTimeout(%q, 1s) = %q, want %q", server.URL+"/hello", string(body), "Hello, World!")
	}
	
	// Test timeout with slow response
	_, err = FetchWithTimeout(server.URL+"/slow", 500*time.Millisecond)
	if err == nil {
		t.Error("FetchWithTimeout with short timeout should return error")
	}
	
	// Test with invalid URL
	_, err = FetchWithTimeout("http://invalid-url-that-does-not-exist.example", 1*time.Second)
	if err == nil {
		t.Error("FetchWithTimeout with invalid URL should return error")
	}
}