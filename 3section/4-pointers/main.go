package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue %+v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	*val = *val * 10 //dereferencing the pointer
	fmt.Printf("modifyValue %+v\n", *val)
}

func main() {

	//age := 10 //
	//agePtr := &age
	//fmt.Printf("age %d\n", &age)
	//fmt.Printf("age addr %d\n", &agePtr)

	num := 10

	modifyValue(num)
	fmt.Println(num)

	modifyPointer(&num)
	fmt.Println(num)

	grade := 50
	gradePtr := &grade
	fmt.Printf("grade: %v\n", gradePtr)        // This is the address of grade
	fmt.Printf("gradePtr: %v\n", &gradePtr)    // This is the address of gradePtr
	fmt.Printf("gradePtr: %v\n", &(*gradePtr)) // dereference
}
