package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Result struct {
	URL      string
	Filename string
	Size     int64
	Duration time.Duration
	Error    error
}

func ConcurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	results := make(chan Result)

	var wg sync.WaitGroup

	limiter := make(chan struct{}, maxConcurrent)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			limiter <- struct{}{}
			defer func() { <-limiter }()

			start := time.Now()
			filename := filepath.Base(url)
			path := filepath.Join(destDir, filename)

			out, err := os.Create(path)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}

			defer out.Close()

			res, err := http.Get(url)

			if err != nil {
				results <- Result{URL: url, Error: err}
			}

			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				results <- Result{URL: url, Error: fmt.Errorf("bad status: %s", res.Status)}
				return
			}

			size, err := io.Copy(out, res.Body)

			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}

			timeSince := time.Since(start)
			results <- Result{
				URL:      url,
				Filename: filename,
				Size:     size,
				Duration: timeSince,
				Error:    nil,
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errors []error
	start := time.Now()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error downloading %s: %s\n", result.URL, result.Error.Error())
			errors = append(errors, result.Error)
		} else {
			totalSize += result.Size
			fmt.Printf("downloaded %s (%d) bytes in %s\n", result.Filename, result.Size, result.Duration)
		}
	}

	startedSince := time.Since(start)
	fmt.Printf("All Downloads completed in %s, Total: %d bytes\n", startedSince, totalSize)

	if len(errors) > 0 {
		return fmt.Errorf("errors downloading %+v", errors)
	}

	return nil
}

func main() {

	urls := []string{
		"https://go.dev/images/go-logo-white.svg",
		//"https://fastly.picsum.photos/id/253/200/300.jpg",
	}

	err := ConcurrentDownloader(urls, "./down", 3)

	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Done")
}
