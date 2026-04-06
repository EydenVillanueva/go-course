package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Write file
	filePath := "./10section/1-files/text.txt"
	data := "Welcome to Go, programming language! Lots of love for Go"
	err := os.WriteFile(filePath, []byte(data), 0644)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done writing")

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	/*
		file2, err := os.Create("file-via-create.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer file2.Close()

		_, err = file2.WriteString("Welcome all Java, Python and Javascript developers")

		if err != nil {
			log.Fatal(err)
		}
	*/

	fileName := "file-via-create.txt"
	printContent(fileName)

	newFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()
	_, _ = newFile.WriteString(fmt.Sprintln("- C"))
	_, _ = newFile.WriteString(fmt.Sprintln("- Ada"))
	_, _ = newFile.WriteString(fmt.Sprintln("- Rust"))
}

func printContent(fileName string) {
	newFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()

	scanner := bufio.NewScanner(newFile)
	lineNum := 1
	for scanner.Scan() {
		lineNum++
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
