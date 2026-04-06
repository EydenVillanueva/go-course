package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text1 := "Hello world! welcome to GO!"

	regGo, err := regexp.Compile(`Go`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Text '%s', matches 'Go': %t\n", text1, regGo.MatchString(text1))

	text2 := "Product Codes: P123, X342, P789"

	rProduct, err := regexp.Compile(`P\d+`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	firstProduct := rProduct.FindString(text2)
	fmt.Println(firstProduct)

	allProducts := rProduct.FindAllString(text2, -1)
	
	fmt.Println(allProducts)

}
