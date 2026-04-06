package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
}

func main() {
	messages := make(chan string)
	users := make(chan user) // Unbuffered channel

	go func() {
		fmt.Println("Sending a message to messages channel")
		messages <- "Hello from messages channel"
	}()

	go func() {
		fmt.Println("Sending a message to messages channel")
		users <- user{
			name: "Gopher",
		}
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("About to get message from channel")

	msg := <-messages

	fmt.Println(msg)

	usr := <-users
	fmt.Println(usr)

}
