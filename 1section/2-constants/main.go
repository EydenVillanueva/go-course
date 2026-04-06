package main

import "fmt"

const HOST string = "localhost"

func main() {
	fmt.Println(HOST)

	const AppName string = "Go"
	fmt.Println(AppName)

	const pi float64 = 3.1415926
	fmt.Println(pi)

	const rate float32 = 5.2
	fmt.Println(rate)

}
