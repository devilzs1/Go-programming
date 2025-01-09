package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Learning working with POST Web requests using Go programming")

    const myURL = "https://localhost:8080/createModule"
	postFORMRequest(myURL)
}

func checkError(err error) {
	if err != nil {
        fmt.Printf("Error making request: %v\n", err)
        panic(err)
    }
}

func postFORMRequest(myURL string) {
	fmt.Println("Making POST request with Form Data to:", myURL)

	data := url.Values{}
	data.Add("title", "Go Programming Series")
	data.Add("module", "Post Form data")
	data.Add("description", "This is a module in the Go Programming Series.")
	data.Add("isLockingEnabled", "true")

	
	response, err := http.PostForm(myURL, data)
	checkError(err)
	
	defer response.Body.Close()

	fmt.Printf("Response status: %d\n", response.StatusCode)
	fmt.Println("Response Content Length:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	var responseBody1 = string(content)
	fmt.Printf("Response body: %s\n", responseBody1)

}