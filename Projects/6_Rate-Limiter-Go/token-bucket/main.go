package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


type Message struct{
	Status string `json:"status"`
	Body string `json:"body"`
}


func endpointHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	message := Message{
		Status: "Successful",
		Body: "Congratulation! You hit the api successfully.",
	}
	json.NewEncoder(w).Encode(&message)
	
}


func main(){
	fmt.Println("In main func of token-bucket")
	http.Handle("/ping", rateLimiter(endpointHandler))
	log.Fatal(http.ListenAndServe(":4000", nil))
}