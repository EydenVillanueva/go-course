package main

import (
	"fmt"
)

type BusinessPerson struct {
	ID   int
	Name string
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

func (e BusinessPerson) String() string {
	return fmt.Sprintf("BusinessPerson[ID: %d, Name: %s]", e.ID, e.Name)
}

func main() {
	jane := BusinessPerson{
		ID:   1,
		Name: "jane",
	}

	fmt.Println(jane)
}
