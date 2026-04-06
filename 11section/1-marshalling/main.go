package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string  `json:"name" xml:"name"`
	Age      int     `json:"age" xml:"age"`
	Phone    string  `json:"phone" xml:"phoneNumber"`
	Password string  `json:"-" xml:"-"`
	IsActive bool    `json:"active" xml:"active"`
	Role     string  `json:"-" xml:"role"`
	Profile  profile `json:"profile" xml:"profile"`
}

type profile struct {
	URL string `json:"url"`
}

var payload = `{
	"name": "Jane",
	"age": 34,
	"phone": "34893483",
	"active": true,
	"password": "12345",
	"profile": {
		"url": "https://ww.jane.co.id"
	}
}`

func main() {

	//marshal
	// jane := user{
	// 	Name:     "Jane",
	// 	Age:      34,
	// 	Phone:    "34893483",
	// 	IsActive: true,
	// }

	// byteSlice, err := json.Marshal(jane)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(byteSlice))

	// unmarshal

	var u user
	err := json.Unmarshal([]byte(payload), &u)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u)
}
