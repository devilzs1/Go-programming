package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
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


func clientRateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler{
	type Client struct{
		limiter *rate.Limiter
		lastSeen time.Time
	}
	var (
		mu sync.Mutex
		clients = make(map[string]*Client)
	)

	go func(){
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3 * time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error getting host & port" , err)
		}
		mu.Lock()

		if _, found := clients[ip]; !found {
			clients[ip] = &Client{limiter : rate.NewLimiter(2,4)}
		}
		clients[ip].lastSeen = time.Now()
		if !clients[ip].limiter.Allow() {

			mu.Unlock()
			msg := Message{
				Status: "Request Failed",
				Body:   "API rate limit exceeded for client : " + ip + " !!!, try again later.",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&msg)
			return
		}
		mu.Unlock()

		next(w,r)
	})
}


func main(){
	fmt.Println("In main func of per-client-rate-limiter")
	http.Handle("/ping", clientRateLimiter(endpointHandler))
	log.Fatal(http.ListenAndServe(":4000", nil))

}