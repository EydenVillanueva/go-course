package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Println("Contact already exists: ", name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}

	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Contact added: %v\n", name)
}

func findContactByName(name string) *Contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}
	return nil
}

func ListContacts() {
	fmt.Println("----- Listing Contacts -----")

	if len(contactList) == 0 {
		return
	}

	for _, contact := range contactList {
		fmt.Printf(
			"id: %d name: %s email: %s phone: %s\n",
			contact.ID, contact.Name, contact.Email, contact.Phone)
	}
}

func main() {
	addContact("Alice Wonderland", "alice@example.com", "111-22222")
	addContact("Bob Builder", "bob@build.com", "222-33333")
	addContact("Charlie Chocolate", "charlie@wonka.com", "333-44444")
	addContact("Dora Explorer", "dora@mapmail.com", "444-55555")
	addContact("Edward Scissorhands", "edward@hedgeart.com", "555-66666")
	addContact("Fiona Shrek", "fiona@swampmail.com", "666-77777")
	addContact("Gandalf Grey", "gandalf@middleearth.me", "777-88888")
	addContact("Homer Simpson", "homer@springfield.net", "888-99999")
	addContact("Indiana Jones", "indy@templemail.org", "999-00000")
	addContact("Jon Snow", "jon@winterfell.north", "000-11111")
	addContact("Katniss Everdeen", "katniss@district12.com", "111-22223")
	addContact("Alice Wonderland", "alice@example.com", "111-22222")

	ListContacts()

	if bob := findContactByName("Bob Builder"); bob != nil {
		fmt.Println("Bob contact found")
		fmt.Println(*bob)
	}

}
