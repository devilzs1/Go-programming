package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Learning working with POST Web requests using Go programming")

    const url = "https://localhost:8080/createSeries"
	postJSONRequest(url)
}

func checkError(err error) {
	if err != nil {
        fmt.Printf("Error making request: %v\n", err)
        panic(err)
    }
}

func postJSONRequest(url string) {
	fmt.Println("Making POST request to:", url)

	requestBody := strings.NewReader(`{
        "title": "Go Programming Series",
        "author": "John Doe",
        "genre": "Technology",
		"description": "This is a series of articles on Go programming language."
    }`)

	response, err := http.Post(url, "application/json", requestBody)
	checkError(err)

	defer response.Body.Close()

	fmt.Printf("Response status: %d\n", response.StatusCode)
	fmt.Println("Response Content Length:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	var responseBody1 = string(content)
	fmt.Printf("Response body: %s\n", responseBody1)


	var responseBody2 strings.Builder
	dataBytes, _ := responseBody2.Write(content)
	fmt.Printf("Response body byte length: %d\n", dataBytes)
	bodyString := responseBody2.String()
	fmt.Println("Response Body:", bodyString)
}