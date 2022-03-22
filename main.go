package main

import (
	"go-rest-api/database"
	"go-rest-api/helpers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	database.Connect()
	router := helpers.NewRouter()
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
