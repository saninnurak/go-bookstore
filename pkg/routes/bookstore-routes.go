package routes

import {
	"github.com/gorilla/mux"
	"github.com/saninnurak/go-bookstore/pkg/controllers"
}

var RegisterBookStoreRoutes = func(router *mux. Router){
	router.HandlerfFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/book/", controllers.GetBook),Methods("GET")
	router.HandlerFunc("/book/{bookId}", controllers.GetBookById),Methods("GET")
	router.HandlerFunc("/book/{bookId}", controllers.UpdateBook),Methods("PUT")
	router.HandlerFunc("/book/{bookId}", controllers.DeleteeBook),Methods("DELETE")
}	
