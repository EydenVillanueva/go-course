package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// Optional to pass numbers
// if you don't pass any param, numbers will be empty slice []
func config(numbers ...int) {
	if len(numbers) > 0 {
		first := numbers[0]
		fmt.Println(first)
	} else {
		fmt.Println("Default number")
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))

	config(5)
	config()
}
