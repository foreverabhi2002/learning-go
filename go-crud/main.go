package main

import (
	"go-crud/config"
	"go-crud/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()
	router := mux.NewRouter()
	routes.RegisterItemRoutes(router)

	log.Println("Server is running on port 5555")
	log.Fatal(http.ListenAndServe(":5555", router))
}
