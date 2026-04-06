package main

import (
	"fmt"
)

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No addresss provided"
	}

	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Printf("Customer ID: %d\n", c.CustomerID)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Email: %s\n", c.Email)
	fmt.Printf("BillingAddress: %s\n", c.BillingAddress.FullAddress())
	fmt.Printf("ShippingAddress: %s\n", c.ShippingAddress.FullAddress())
}

func main() {

	fmt.Println("---- Composition ----")
	c1 := Customer{
		CustomerID: 1001,
		Name:       "Gadget Corp",
		Email:      "sales@gadgetcorp.com",
		BillingAddress: Address{
			Street:  "123 tech road",
			City:    "Innovateville",
			State:   "CA",
			ZipCode: "90381",
		},
		ShippingAddress: Address{
			Street:  "472 factory lane",
			City:    "Manucity",
			State:   "CA",
			ZipCode: "90381",
		},
	}
	c1.PrintDetails()
}
