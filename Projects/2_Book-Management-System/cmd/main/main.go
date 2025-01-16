package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devilzs1/book-store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Book Management Store - Build using Go & MySQL")

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":4000", r))
	
}