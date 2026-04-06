package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstName, lastName, position string, salary int, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

func main() {

	jane := Employee{
		ID:        1,
		FirstName: "jane",
		LastName:  "Smith",
		Position:  "Employee",
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	//fmt.Printf("%+v\n", jane)
	fmt.Println(jane.ID)
	fmt.Println(jane.FirstName)

	joe := NewEmployee(1, "john", "doe", "support", 10, true)
	fmt.Println(joe.Salary)
}
