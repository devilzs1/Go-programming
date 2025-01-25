package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/didip/tollbooth/v8"
	"github.com/didip/tollbooth/v8/limiter"
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
	fmt.Println("In main func of tollbooth based rate limiter")
	msg := Message{
		Status: "Request Failed",
		Body:   "API rate limit exceeded!!!, try again later.",
	}
	jsonMsg, _ := json.Marshal(msg)

	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetMessageContentType("application/json")
	lmt.SetMessage(string(jsonMsg))

	lmt.SetIPLookup(limiter.IPLookup{
		Name:           "X-Real-IP",
		IndexFromRight: 0,
	})

	http.Handle("/ping", tollbooth.HTTPMiddleware(lmt)(http.HandlerFunc(endpointHandler)))

	log.Fatal(http.ListenAndServe(":4000", nil))
}