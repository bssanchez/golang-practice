package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

// Get makes an HTTP GET request and returns the response body
func Get(url string) ([]byte, error) {
	// TODO: Implement this function
	return nil, nil
}

// GetJSON makes an HTTP GET request and deserializes the JSON response
func GetJSON(url string, v interface{}) error {
	// TODO: Implement this function
	return nil
}

// Post makes an HTTP POST request and returns the response body
func Post(url string, contentType string, body []byte) ([]byte, error) {
	// TODO: Implement this function
	return nil, nil
}

// DownloadFile downloads a file from a URL and saves it to the specified path
func DownloadFile(url, filepath string) error {
	// TODO: Implement this function
	return nil
}

// FetchWithTimeout makes an HTTP GET request with a specified timeout
func FetchWithTimeout(url string, timeout time.Duration) ([]byte, error) {
	// TODO: Implement this function
	return nil, nil
}