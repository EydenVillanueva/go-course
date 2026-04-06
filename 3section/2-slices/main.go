package main

import "fmt"

func main() {
	names := []string{"Alice", "John", "Mike"}
	fmt.Println(names)

	// If you know the capacity ahead, declare the capacity in advance
	//             lenght,   capacity
	// make([]int, 3,        10)
	items := make([]int, 3, 10)
	fmt.Printf("%+v\n", items)

	newitems := make([]int, 3, 5)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", newitems, len(newitems), cap(newitems))

	newitems = append(newitems, 1)
	newitems = append(newitems, 2)
	newitems = append(newitems, 3)
	newitems = append(newitems, 4)

	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", newitems, len(newitems), cap(newitems))

	fmt.Printf("%+v\n", newitems[3:7])
}
