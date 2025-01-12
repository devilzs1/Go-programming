package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

)

// generally preferred as a pointer because we use them in multiple go routines
var wg sync.WaitGroup
var mut sync.Mutex  

var signals = []string{"test"}

func main() {
	fmt.Println("Learning about Concurrency and Go Routines in Go Programming")

	websiteList := []string{
		"https://go.dev", "https://github.com/devilzs1", "https://google.com", "https://github.com",
	}

	for _, en := range websiteList {
		go getStatusCode(en)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}


func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for endpoint %s\n", res.StatusCode, endpoint)

}