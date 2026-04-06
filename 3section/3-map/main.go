package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"Alice": 90,
		"James": 85,
		"Dan":   60,
	}

	fmt.Printf("%+v\n", studentGrades)

	studentGrades["Alice"] = 95
	fmt.Printf("%+v\n", studentGrades)

	// if ok = true, The value exist
	// if ok = false, the value doesnt exist
	alice, ok := studentGrades["Alice"]

	if ok {
		fmt.Println("Alice", alice)
	}

	key := "Dan"
	if _, ok := studentGrades[key]; ok {
		fmt.Printf("%s, %+v\n", key, studentGrades[key])
	}

	key = "Bob"

	// Bob doesnt exist
	if _, ok := studentGrades[key]; ok {
		fmt.Printf("%s, %+v\n", studentGrades[key])
	}

	delete(studentGrades, "Alice")
	fmt.Printf("%+v\n", studentGrades)

	// Initalize map, two ways:
	// 1 ) configs := make(map[string]int)
	// 2 ) configs := map[string]int{}Ø
	configs := map[string]int{}
	fmt.Printf("%+v %T\n", configs, configs)

	if configs == nil {
		fmt.Println("Config is nil")
	}
}
