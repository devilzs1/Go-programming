package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Learning Docker integration using Go Programming")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world! Contacting from docker container api.....")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}


// Build docker image : docker build -t go-docker-app .
// Run the docker container : docker run -p 8080:8080 -it go-docker-app
