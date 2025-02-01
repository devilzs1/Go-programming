package utils

import (
	"log"
	"os"
)



func HandleError(err error) {
	if err != nil {
		log.Fatalf("Error : ", err)
		os.Exit(1)
	}
}