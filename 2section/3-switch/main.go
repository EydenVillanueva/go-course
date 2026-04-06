package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Sunday"

	fmt.Println("Today is", day)

	switch day {

	case "Sunday", "Saturday":
		fmt.Println("Weekend! not work")
	case "Monday", "Tuesday":
		fmt.Println("Work days. Lots of meetings")
	default:
		fmt.Println("Mid-Week")

	}

	switch hour := time.Now().Hour(); {

	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")

	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("Twice %v is %v\n", v, day)
		case string:
			fmt.Println("String %s\n", v)
		case bool:
			fmt.Println("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type : %T\n", v)

		}
	}

	checkType(21)
	checkType("string")
	checkType(true)
	checkType(3.23435)
}
