package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 3)

	fmt.Println("Sending messages to buffered channel")

	messages <- "Hello from messages channel"
	messages <- "Hello from messages channel"
	messages <- "Hello from messages channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
