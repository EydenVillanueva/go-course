package main

import "fmt"

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

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", ci.Email, ci.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)

	fmt.Printf("Location: %s\n", c.FullAddress())
	fmt.Printf("Steet (promoted): %s\n", c.Street)

	fmt.Printf("Email (promoted): %s\n", c.Email)
	fmt.Printf("BusinessType: %s\n", c.BusinessType)
	fmt.Printf("Phone (promoted): %s\n", c.Phone)
}

func main() {

	fmt.Println(" ---- Struct Embedding ----")

	comp := Company{
		Name: "Apple",
		Address: Address{
			Street:  "street",
			City:    "city",
			State:   "state",
			ZipCode: "230924",
		},
		ContactInfo: ContactInfo{
			Email: "email@email.com",
			Phone: "0293942",
		},
		BusinessType: "Technology",
	}

	comp.GetProfile()

}
