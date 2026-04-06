package main

import (
	"errors"
	"fmt"
	"strings"
)

// By convention error is the last to be returned
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	return
}

func main() {
	value, err := divide(0, 4)
	if err != nil {
		if err.Error() == "a is too large" {
			fmt.Println("do something else")
		}
	} else {
		fmt.Println(value)
	}

	firstName, lastName := splitName("Eyden Villanueva")
	fmt.Println(firstName, lastName)
}
