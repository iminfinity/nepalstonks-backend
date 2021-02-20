package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iminfinity/nepalstonks/models"
	"gopkg.in/mgo.v2/bson"
)

// GetStockData func
func GetStockData(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	params := mux.Vars(r)
	stock := params["stock"]
	var stockData models.StockDataList

	err = stocksCollection.FindOne(ctx, bson.M{"stockName": stock}).Decode(&stockData)
	if err != nil {
		http.Error(rw, "Stock not available", http.StatusInternalServerError)
		fmt.Println("Get request failed")
		return
	}

	json.NewEncoder(rw).Encode(&stockData)
	fmt.Println("Get request successful")
}

// Give func
func Give(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Working")
}
