package routes

import (
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"github.com/devilzs1/book-store/pkg/controllers"

)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}