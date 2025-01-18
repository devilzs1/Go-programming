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
)

type response struct {
	Id      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file : ", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("Error connecting to db : ", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to postgres sql")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var newStock model.Stock

	err := json.NewDecoder(r.Body).Decode(&newStock)
	if err != nil {
		log.Fatal("Error decoding request body : ", err)
	}
	insertId := insertStock(newStock)

	res := response{
		Id:      insertId,
		Message: "Stock created successfully.",
	}
	json.NewEncoder(w).Encode(res)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()
	if err != nil {
		log.Fatal("Error fetching all the stocks : ", err)
	}
	json.NewEncoder(w).Encode(stocks)
}

func GetStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId, err := strconv.Atoi(params["stockId"])
	if err != nil {
		log.Fatal("Unable to convert Id to int : ", err)
	}
	stock, err := getStock(int64(stockId))
	if err != nil {
		log.Fatal("Error finding stock for the given Id : ", err)
	}
	json.NewEncoder(w).Encode(stock)

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId, err := strconv.Atoi(params["stockId"])
	if err != nil {
		log.Fatal("Unable to convert Id to int : ", err)
	}
	deletedRows := deleteStock(int64(stockId))
	msg := fmt.Sprintf("Stock updated successfully. Total rows affected : %v", deletedRows)
	res := response{
		Id:      int64(stockId),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId, err := strconv.Atoi(params["stockId"])
	if err != nil {
		log.Fatal("Unable to convert Id to int : ", err)
	}
	var stock model.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode request body : ", err)
	}
	updatedRows := updateStock(int64(stockId), stock)
	msg := fmt.Sprintf("Stock updated successfully. Total rows affected : %v", updatedRows)
	res := response{
		Id:      int64(stockId),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func insertStock(stock model.Stock) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3, $4) RETURNING stockId`
	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company, stock.Quantity).Scan(&id)
	if err != nil {
		log.Fatal("Error inserting stock into db : ", err)
	}
	fmt.Println("Stock inserted successfully : ", id)
	return id
}

func updateStock(id int64, stock model.Stock) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `UPDATE stocks SET stockName=$2, price=$3, company=$4, quantity=$5 WHERE stockId=$1`
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company, stock.Quantity)
	if err != nil {
		log.Fatal("Error updating the stock : ", err)
	}
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("Stock updated successfully. Total rows affected : ", rowsAffected)
	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM stocks WHERE stockId=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal("Error deleting the stock : ", err)
	}
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("Stock deleted successfully. Total rows affected : ", rowsAffected)
	return rowsAffected
}

func getStock(id int64) (model.Stock, error) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `SELECT * FROM stocks WHERE stockId = $1`
	var stock model.Stock
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.Id, &stock.Name, &stock.Price, &stock.Company, &stock.Quantity)
	switch err {
	case sql.ErrNoRows:
		{
			fmt.Println("No rows returned!")
			return stock, nil
		}
	case nil:
		{
			return stock, nil
		}
	default:
		{
			fmt.Println("Unable to scan rows!")
		}
	}
	return stock, err
}

func getAllStocks() ([]model.Stock, error) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `SELECT * FROM stocks`
	var stocks []model.Stock
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("Error fetching stocks : ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var stock model.Stock
		err := rows.Scan(&stock.Id, &stock.Name, &stock.Price, &stock.Company, &stock.Quantity)
		if err != nil {
			log.Fatal("Error scanning rows : ", err)
		}
		stocks = append(stocks, stock)
	}
	fmt.Println("Stocks fetched successfully")
	return stocks, err
}
