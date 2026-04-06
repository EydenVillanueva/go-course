package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(" ------ Advanced Slicing Operations --------")
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("original: %v, len %d, cap: %d\n", original, len(original), cap(original))

	s1 := original[2:5] // low:high (excluding high)
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))

	s2 := original[:4] // low:high (high is not included)
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s3 := original[6:]
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))

	s4 := original[:]
	fmt.Printf("s4: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))

	slices.Contains(s4, 8)
	slices.Insert(s4, 10000)

	fmt.Printf("s4: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))
}
