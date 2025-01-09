package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Learning how to make GET Request using Go programming language")

	const url = "http://github.com/devilzs1"
	getRequest(url)
}

func checkError(err error) {
	if err != nil {
        fmt.Printf("Error making request: %v\n", err)
        panic(err)
    }
}

func getRequest(url string) {
	fmt.Println("Making GET request to:", url)

	response, err := http.Get(url)
	checkError(err)

	defer response.Body.Close()

	fmt.Printf("Response status: %d\n", response.StatusCode)
	fmt.Println("Response Content Length:", response.ContentLength)

	var responseBody strings.Builder
	content, _ := io.ReadAll(response.Body)
	dataBytes, _ := responseBody.Write(content)

	fmt.Printf("Response body byte length: %d\n", dataBytes)

	// bodyString := responseBody.String()
	// fmt.Println("Response Body:", bodyString)
}