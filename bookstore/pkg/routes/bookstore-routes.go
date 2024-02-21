package routes

import (
	"github.com/gorilla/mux"
	"github.com/toliveiradev/learning-go/bookstore/pkg/controller"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE")
}
