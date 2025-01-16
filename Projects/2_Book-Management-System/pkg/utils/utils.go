package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}){
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			log.Println("Error unmarshaling JSON into struct:", err)
			return 
		}
	} else {
		log.Println("Error reading request body:", err)
	}
}
