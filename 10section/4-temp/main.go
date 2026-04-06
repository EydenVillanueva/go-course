package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tempFile, err := os.CreateTemp("", "logs.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Removing tempFile", tempFile.Name())
		_ = os.Remove(tempFile.Name())
	}()

	_, err = tempFile.WriteString("Hello World\n")

	if err != nil {
		log.Fatal(err)
		tempFile.Close()
		return
	}
}
