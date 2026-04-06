package main

import (
	"fmt"
	"time"
)

func sayHello(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("sayHello", message)
}

func main() {
	fmt.Println("Hello from Main() Goroutine")

	//
	go sayHello("Hello World", time.Second)
	go sayHello("Hello World 2", time.Second)
	go sayHello("Hello World 3", time.Second)

	fmt.Println("last message from Main() Goroutine")
	time.Sleep(2 * time.Second)
}
