package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	data := "Welcome to the wonderful world of Go!"

	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(encoded)

	encodedStr := "V2VsY29tZSB0byB0aGUgd29uZGVyZnVsIHdvcmxkIG9mIEdvIQ=="
	decoded, err := base64.StdEncoding.DecodeString(encodedStr)

	if err != nil {
		log.Fatal(err)
	}

	if string(decoded) != data {
		log.Fatalf("Decoded string does not match encoded data")
	}

	rawData := []byte{0xDE, 0xAD, 0xEF, 0xCA, 0xBA}
	binaryCodedToString := base64.StdEncoding.EncodeToString(rawData)

	fmt.Println(string(binaryCodedToString))

	base64Str := "3q3vyro="
	decoded, err = base64.RawStdEncoding.DecodeString(base64Str)

	if err != nil {
		log.Fatal(err)
	}
}
