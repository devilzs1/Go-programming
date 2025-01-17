package router

import (
	controller "github.com/devilzs1/stocks/controllers"
	"github.com/gorilla/mux"
)



func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/stocks", controller.GetAllStocks).Methods("GET")
	router.HandleFunc("/stock/new", controller.CreateStock).Methods("POST")
	router.HandleFunc("/stock/{stockId}", controller.UpdateStock).Methods("PUT")
	router.HandleFunc("/stock/{stockId}", controller.DeleteStock).Methods("DELETE")
	router.HandleFunc("/stock/{stockId}", controller.GetStockById).Methods("GET")

	return router
}