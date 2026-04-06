package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func intSequence() func() int {
	i := 1
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(factorial(5))

	nextInt := intSequence()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger := func(msg string) {
		fmt.Println(msg)
	}

	logger("hello world")
}
