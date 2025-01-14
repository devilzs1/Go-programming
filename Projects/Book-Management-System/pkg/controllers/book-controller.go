package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/devilzs1/book-store/pkg/models"
	"github.com/gorilla/mux"
)


var NewBook models.Book


func GetAllBook(w http.ResponseWriter, r * http.Request){
	newBooks := models.GetAllBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter, r * http.Request){
	
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		panic(err)
	}
	bookDetails, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

