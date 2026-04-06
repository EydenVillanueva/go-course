package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	URL      string
	Status   string
	Duration time.Duration
	Err      error
}

func checkWebsite(url string, ch chan<- Result) {
	start := time.Now()

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		ch <- Result{URL: url, Status: "DOWN", Err: err}
		return
	}

	defer resp.Body.Close()

	ch <- Result{
		URL:      url,
		Status:   "OK",
		Duration: duration,
	}
}

func main() {
	websites := []string{
		"https://beta.getobok.com",
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://doesnotexist.tld",
	}

	results := make(chan Result)

	for _, site := range websites {
		go checkWebsite(site, results)
	}

	for range websites {
		result := <-results

		if result.Err != nil {
			fmt.Printf("[DOWN] %s (%s)\n", result.URL, result.Err)
		} else {
			fmt.Printf("[OK] %s (%v)\n", result.URL, result.Duration)
		}
	}
}
