package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed hello.txt
var data string

//go:embed public
var public embed.FS

func main() {
	fmt.Println(data)

	data, err := public.ReadFile("public/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	// After building main.go with go build
	// All the assets (In this case the public directory)
	// Will be embedded into the main executable binary
	fmt.Println(string(data))
}
