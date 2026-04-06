package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abc"
	s2 := strings.Clone(s1)

	fmt.Println(s2)

	b := strings.Builder{}
	b.WriteString("Here is an example")
	fmt.Println(b.String())

	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s2))
	fmt.Println(strings.ToTitle(s2))

	fmt.Println(strings.TrimSpace("     test    ss    "))
}
