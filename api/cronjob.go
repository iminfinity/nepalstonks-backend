package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/iminfinity/nepalstonks/models"
	"github.com/iminfinity/nepalstonks/utils"
	"gopkg.in/mgo.v2/bson"
)

func updateSingleStock(stockFromHerocu models.ResponsedDataFromHerocu) {
	stockSymbol := stockFromHerocu.StockSymbol

	var stock models.StockDataList

	err = stocksCollection.FindOne(ctx, bson.M{"stockName": stockSymbol}).Decode(&stock)
	if err != nil {
		fmt.Println("Update failed")
		return
	}

	details := stock.StockData
	details.MaxPrice = utils.PustItemToFirst(details.MaxPrice, stockFromHerocu.MaxPrice)
	details.MinPrice = utils.PustItemToFirst(details.MinPrice, stockFromHerocu.MinPrice)
	details.ClosingPrice = utils.PustItemToFirst(details.ClosingPrice, stockFromHerocu.ClosingPrice)

	date := strings.Split(time.Now().String(), " ")

	details.Date = utils.PustItemToFirst(details.Date, date[0])

	stock.StockData = details
	_, err = stocksCollection.UpdateOne(ctx, bson.M{"stockName": stockSymbol}, bson.M{"$set": stock})
	if err != nil {
		fmt.Println("Update failed")
		IfNewThenAdd(stockSymbol)
		return
	}

}

func updateAllData() bool {
	response, err := http.Get("https://nepstockapi.herokuapp.com/")
	if err != nil {
		fmt.Println("Error")
		return false
	}
	data, err := ioutil.ReadAll(response.Body)
	var rawData []models.ResponsedDataFromHerocu

	err = json.Unmarshal([]byte(data), &rawData)
	if err != nil {
		fmt.Println("Update failed")
		return false
	}
	for _, currentStock := range rawData {
		updateSingleStock(currentStock)
	}

	return true
}

// UpdateEveryDay updates stock data everyday 3:10pm
func UpdateEveryDay(rw http.ResponseWriter, r *http.Request) {
	TOKEN := os.Getenv("STONKS_TOKEN")

	rw.Header().Add("content-type", "applicaiton/json")
	params := mux.Vars(r)
	token := params["token"]
	if token != TOKEN {
		http.Error(rw, "Not Authorised", http.StatusForbidden)
		fmt.Println("Unauthorised access denied")
		return
	}
	success := updateAllData()

	if success {
		fmt.Println("Update successfull")
	} else {
		fmt.Println("Update Failed")
	}
}

// RemoveDuplicate func
func RemoveDuplicate(rw http.ResponseWriter, r *http.Request) {
	var stockList models.StockList

	err = stocksListCollection.FindOne(ctx, bson.M{"bodge": "giveData"}).Decode(&stockList)
	if err != nil {
		http.Error(rw, "Failed getting stock list", http.StatusInternalServerError)
		fmt.Println("Failed getting stock list")
		return
	}

	for _, currentStock := range stockList.List {
		var currentStockData models.StockDataList

		err = stocksCollection.FindOne(ctx, bson.M{"stockName": currentStock}).Decode(&currentStockData)
		if err != nil {
			fmt.Println("Removing duplicate func failed")
			return
		}

		details := currentStockData.StockData

		if len(details.Date) <= 2 {
			continue
		}
		if details.Date[0] != details.Date[1] {
			continue
		}

		details.Date = details.Date[1:]
		details.MaxPrice = details.MaxPrice[1:]
		details.MinPrice = details.MinPrice[1:]
		details.ClosingPrice = details.ClosingPrice[1:]

		currentStockData.StockData = details
		_, err = stocksCollection.UpdateOne(ctx, bson.M{"stockName": currentStock}, bson.M{"$set": currentStockData})
		if err != nil {
			fmt.Println("Removing duplicate func failed")
			return
		}
	}
	fmt.Println("Remove Duplicate ran successfully")
}
