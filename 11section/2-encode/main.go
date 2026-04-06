package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phoneNumber"`
	Password string `json:"-" xml:"-"`
	IsActive bool   `json:"active" xml:"active"`
}

func main() {

	u := user{
		Name:     "John Smith",
		Phone:    "1238429",
		Age:      42,
		IsActive: true,
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())

}
