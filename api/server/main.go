package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// HandleDuress sends an OK header from the server
func HandleDuress(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	vars := mux.Vars(req)
	fmt.Printf("duress check on [%s_%s] [%v]\n", vars["clientID"], vars["duressCode"], time.Since(start))
	w.WriteHeader(200)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/client/{clientID}/code/{duressCode}", HandleDuress)
	log.Fatal(http.ListenAndServe(":8000", r))
}
