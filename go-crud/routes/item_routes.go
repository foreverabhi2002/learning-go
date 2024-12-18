package routes

import (
	"go-crud/controllers"

	"github.com/gorilla/mux"
)

func RegisterItemRoutes(router *mux.Router) {
	router.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/items", controllers.GetItems).Methods("GET")
}
