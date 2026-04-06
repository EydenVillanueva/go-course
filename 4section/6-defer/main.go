package main

import (
	"fmt"
)

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	// When you defer a function, gets executed immediately after the function returns
	defer fmt.Println("Function simpleDefer: deferred")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
}

func lifoSimpleDefer() {
	fmt.Println("Function simpleDefer:Start")
	defer fmt.Println("First: deferred")
	defer fmt.Println("Second: deferred")
	fmt.Println("Function lifoSimpleDefer: Middle")
}

func main() {
	fmt.Println("Function simpleDefer:Start")
	defer lifoSimpleDefer()

	//file, err := os.Create("./defer.txt")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//defer file.Close()

}
