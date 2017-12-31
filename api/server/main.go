package main

import (
	"log"
	"net/http"

	"github.com/alxunger/go-duress/api/db"
	"github.com/alxunger/go-duress/api/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/client/{clientID}/code/{duressCode}", handlers.HandleDuress)
	r.HandleFunc("/countCodes", handlers.HandleCount)
	log.Fatal(http.ListenAndServe(":8000", r))
}
