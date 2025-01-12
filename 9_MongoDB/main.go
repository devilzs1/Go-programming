package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devilzs1/mongodb-go/router"
)

func main() {
	fmt.Println("Learning integrating MongoDB using Go Programming")

	r := router.Router()
	fmt.Println("Server is starting .....")
	fmt.Println("Server listening on port : ", 4000)
	log.Fatal(http.ListenAndServe(":4000", r))
}