package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning working with Web requests URLs using Go programming")

    const myUrl = "https://github.com:8091/devilzs1?repo=overview&owner=devilzs"
	fmt.Println("URL:", myUrl)

	parts, _ := url.Parse(myUrl)

	fmt.Println("URL scheme:", parts.Scheme)
	fmt.Println("URL host:", parts.Host)
	fmt.Println("URL path:", parts.Path)
	fmt.Println("URL port:", parts.Port())
	fmt.Println("URL query:", parts.RawQuery)

	qParams := parts.Query()
	fmt.Printf("URL Query params: %v & type of queryParams is %T\n", qParams, qParams)
	// repo := qParams["repo"]
	// repo := qParams.Get("repo")
	// fmt.Printf("Value of repo query param: %s\n", repo)

	for key, val := range qParams {
		fmt.Printf("Key: %s, Value: %s\n", key, val)
	}


	partsOfURL := &url.URL{
		Scheme:   "https",
        Host:     "github.com",
        Path:     "/devilzs1/",
	}

	newURL := partsOfURL.String()
	fmt.Println("New URL: ", newURL)
}


func checkError(err error) {
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}
}