package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/alxunger/go-duress/api/db"
	"github.com/gorilla/mux"
)

func HandleDuress(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	vars := mux.Vars(req)
	fmt.Printf("duress check on [%s_%s] [%v]\n", vars["clientID"], vars["duressCode"], time.Since(start))
	w.WriteHeader(200)
}

func HandleCount(w http.ResponseWriter, req *http.Request) {
	totalCodes := db.TotalCodes()
	w.Write([]byte(strconv.Itoa(int(totalCodes))))
}
