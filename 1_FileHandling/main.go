package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Learning handling files...")

	txt := "Hello world! I am learning file handling using Go programming"

	file, err := os.Create("./sample.txt")
	checkError(err)

	defer file.Close()

	length, err := io.WriteString(file, txt)
	checkError(err)

	fmt.Printf("Successfully wrote %d bytes to the file.\n", length)

	readFile("./sample.txt")
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Printf("File content: %s\n", string(data))
}


func checkError(err error) {
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}
}
