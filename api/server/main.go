package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/alxunger/go-duress/api/db"
	"github.com/gorilla/mux"
)

// TODO export handlers to its own package

// HandleDuress sends an OK header from the server
func HandleDuress(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	vars := mux.Vars(req)
	fmt.Printf("duress check on [%s_%s] [%v]\n", vars["clientID"], vars["duressCode"], time.Since(start))
	w.WriteHeader(200)
}

// HandleCount counts the total number of codes stored
func HandleCount(w http.ResponseWriter, req *http.Request) {
	totalCodes := db.CountCodes()
	w.Write([]byte(strconv.Itoa(totalCodes)))
}

func main() {
	db.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/client/{clientID}/code/{duressCode}", HandleDuress)
	r.HandleFunc("/countCodes", HandleCount)
	log.Fatal(http.ListenAndServe(":8000", r))
}
