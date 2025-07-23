# Exercise 8: Concurrent Downloader

## Description

Implement a concurrent download system that can download multiple files simultaneously with concurrency limits.

## Requirements

1. Implement the following structures and functions in the `downloader.go` file:

   - Struct `DownloadResult` with fields for URL, local path, error, and download time
   - Struct `Downloader` with field for maximum number of concurrent downloads
   - Method `NewDownloader(maxConcurrent int) *Downloader` - Constructor
   - Method `Download(url, filepath string) error` - Downloads a single file
   - Method `DownloadMany(urls []string, directory string) []DownloadResult` - Downloads multiple files
   - Method `DownloadWithProgress(url, filepath string, progressChan chan<- int) error` - Downloads with progress notification

2. Considerations:
   - Use goroutines for concurrent downloads
   - Limit the number of simultaneous downloads according to `maxConcurrent`
   - Implement a mechanism to report download progress
   - Handle errors and timeouts correctly
   - Ensure resources are properly released

## Tests

Run `go test` to verify your implementation.

## Note

The tests use a mock HTTP server to avoid external dependencies.
