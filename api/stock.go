package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/iminfinity/nepalstonks/models"
	"github.com/iminfinity/nepalstonks/utils"
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
	var wg sync.WaitGroup

	wg.Add(1)
	go utils.Reverse(&wg, stockData.StockData.Date)
	wg.Add(1)
	go utils.Reverse(&wg, stockData.StockData.MaxPrice)
	wg.Add(1)
	go utils.Reverse(&wg, stockData.StockData.MinPrice)
	wg.Add(1)
	go utils.Reverse(&wg, stockData.StockData.ClosingPrice)

	wg.Wait()
	json.NewEncoder(rw).Encode(&stockData)
	fmt.Println("Get request successful")
}

// Give func
func Give(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Working")
}
