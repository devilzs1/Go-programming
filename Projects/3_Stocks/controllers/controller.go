package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	model "github.com/devilzs1/stocks/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/genproto/googleapis/cloud/aiplatform/v1/schema/predict/params"
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
	var newStock model.Stock

	err := json.NewDecoder(r.Body).Decode(&newStock)
	if err != nil{
		log.Fatal("Error decoding request body : ", err)
	}
	insertId := insertStock(newStock)

	res := response{
		Id: insertId,
		Message: "Stock created successfully.",
	}
	json.NewEncoder(w).Encode(res)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request){
	stocks, err := getAllStocks()
	if err != nil{
		log.Fatal("Error fetching all the stocks : ", err)
	}
	json.NewEncoder(w).Encode(stocks)
}

func GetStockById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	stockId, err := strconv.Atoi(params["stockId"])
	if err != nil{
		log.Fatal("Unable to convert Id to int : ", err)
	}
	stock, err := getStock(stockId)
	if err != nil{
		log.Fatal("Error finding stock for the given Id : ", err)
	}
	json.NewEncoder(w).Encode(stock)

}

func DeleteStock(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	stockId, err := strconv.Atoi(params["stockId"])
	if err != nil{
		log.Fatal("Unable to convert Id to int : ", err)
	}
	deletedRows := deleteStock(stockId)
	msg := fmt.Sprintf("Stock updated successfully. Total rows affected : %v", deletedRows)
	res := response{
		Id: int64(stockId),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateStock(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	stockId, err := strconv.Atoi(params["stockId"])
	if err != nil{
		log.Fatal("Unable to convert Id to int : ", err)
	}
	var stock model.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil{
		log.Fatal("Unable to decode request body : ", err)
	}
	updatedRows := updateStock(int64(stockId), stock)
	msg := fmt.Sprintf("Stock updated successfully. Total rows affected : %v", updatedRows)
	res := response{
		Id: int64(stockId),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}