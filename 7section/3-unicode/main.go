package main

import "fmt"

func main() {
	username := "testñ"

	fmt.Println(len(username)) // This is counting bytes not characters
}
