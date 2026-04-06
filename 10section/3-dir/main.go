package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	dir := "Downloads/static/images"
	if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll(filepath.Clean("Downloads")); err != nil {
		log.Fatal(err)
	}

	tempDir, err := os.MkdirTemp("", "my_app_logs")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Removing file", tempDir)
		_ = os.Remove(tempDir)
	}()
}
