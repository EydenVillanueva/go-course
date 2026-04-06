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

// FullName A value receiver (Copy) It doesn't affect the original reference
func (e *Employee) FullName() string {
	if e == nil {
		return ""
	}
	return e.FirstName + " " + e.LastName
}

func (e *Employee) Deactivate() {
	e.IsActive = false
}

func (e *Employee) SetJoinedAt(t time.Time) {
	e.JoinedAt = t
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

func deactivate(e *Employee) {
	e.IsActive = false
}

func main() {
	jane := Employee{
		ID:        1,
		FirstName: "jane",
		LastName:  "Smith",
		Position:  "Employee",
		IsActive:  true,
		//JoinedAt:  time.Now(),
	}

	//deactivate(&jane)
	jane.Deactivate()
	fmt.Printf("%+v\n", jane)

	jane.SetJoinedAt(time.Now().Add(100000000 * time.Minute))
	fmt.Printf("%+v\n", jane)

	joe := &Employee{}
	joe = nil      // demo to show panic
	joe.FullName() // Bad idea
}
