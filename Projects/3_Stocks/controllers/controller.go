package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)



type response struct{
	Id int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB{
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading env file : ", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil{
		log.Fatal("Error connecting to db : ", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Successfully connected to postgres sql")
	return db
}


func CreateStock(w http.ResponseWriter, r *http.Request) {

}

func GetAllStocks(w http.ResponseWriter, r *http.Request){

}

func GetStockById(w http.ResponseWriter, r *http.Request){

}

func DeleteStock(w http.ResponseWriter, r *http.Request){

}

func UpdateStock(w http.ResponseWriter, r *http.Request){
	
}