package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devilzs1/stocks/router"
)



func main(){
	fmt.Println("Stocks api built using Postres SQL & Go programming")
	router := router.Router()

	fmt.Println("Starting server on port 4000....")
	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Server listening on localhost:4000.....")
}