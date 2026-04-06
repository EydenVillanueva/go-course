package main

import "fmt"

type Number interface {
	int | float32 | float64 | string
}

func Sum[T Number](numbers ...T) T {
	var total T

	for _, n := range numbers {
		total += n
	}

	return total
}

func main() {

	//grades := []int{90, 85}
	//people := []string{"jane", "John", "Mark"}
	//
	//fmt.Println(len(grades), len(people))

	v := Sum(10, 20, 30, 3.3)
	c := Sum("Jane", "Mark")

	fmt.Printf("%T\n", v)
	fmt.Printf("%T\n", c)
}
