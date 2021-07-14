package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	router := NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"Access-Control-Allow-Origin", "*"})
	//ローカル環境 (http)
	//log.Fatal(http.ListenAndServe(":8081", handlers.CORS(originsOk, headersOk)(router)))
	//本番環境 (https)
	log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", handlers.CORS(originsOk, headersOk)(router)))
}
