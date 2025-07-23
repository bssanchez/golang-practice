# Exercise 7: HTTP Client

## Description
Implement a simple HTTP client that can make requests and process responses.

## Requirements
1. Implement the following functions in the `httpclient.go` file:
   - `Get(url string) ([]byte, error)` - Makes a GET request and returns the response body
   - `GetJSON(url string, v interface{}) error` - Makes a GET request and deserializes the JSON response
   - `Post(url string, contentType string, body []byte) ([]byte, error)` - Makes a POST request
   - `DownloadFile(url, filepath string) error` - Downloads a file from a URL
   - `FetchWithTimeout(url string, timeout time.Duration) ([]byte, error)` - Makes a GET request with timeout

2. Considerations:
   - Handle HTTP errors correctly (status codes, timeouts, etc.)
   - For `GetJSON`, use the standard library `encoding/json`
   - For `DownloadFile`, write the content to the file efficiently
   - Implement a mechanism to cancel requests that take too long

## Tests
Run `go test` to verify your implementation.

## Note
The tests use a mock HTTP server to avoid external dependencies.