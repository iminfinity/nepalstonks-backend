package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iminfinity/nepalstonks/api"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/stock-data/{stock}", api.GetStockData).Methods("GET")

	router.HandleFunc("/update-every-day/{token}", api.UpdateEveryDay).Methods("GET")

	router.HandleFunc("/give", api.Give).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Printf("$PORT not set")
	}

	http.ListenAndServe(":"+port, router)

}
