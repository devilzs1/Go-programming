package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	fmt.Println("Learning working with Web requests using Go programming")

	const url = "https://github.com/devilzs1"

	res, err := http.Get(url)
	checkError(err)

	defer res.Body.Close()

	fmt.Printf("Response is of type : %T\n", res)
	fmt.Println("Response status: ", res.StatusCode)
	fmt.Println("Response headers: ", res.Header)

	dataBytes, err := io.ReadAll(res.Body)
	checkError(err)
	content := string(dataBytes)
	fmt.Printf("Response body length: %d\n", len(content))
	// fmt.Println("Response body: ", content)

}

func checkError(err error) {
	if err != nil {
        fmt.Printf("Error making HTTP request: %v\n", err)
        panic(err)
    }
}