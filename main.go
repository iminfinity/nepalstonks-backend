package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iminfinity/nepalstonks/api"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/stock-data/{stock}", api.GetStockData).Methods("GET")

	router.HandleFunc("/api/update-every-day/{token}", api.UpdateEveryDay).Methods("GET")

	router.HandleFunc("/api/stock-list", api.GetStockList).Methods("GET")

	router.HandleFunc("/api/remove-duplicate", api.RemoveDuplicate).Methods("GET")

	router.HandleFunc("/api/wake-up", api.WakeUp).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Printf("$PORT not set")
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://nepalstonks.com", "195.201.26.157", "116.203.134.67", "116.203.129.16"},
	})

	http.ListenAndServe(":"+port, corsHandler.Handler(router))

}
