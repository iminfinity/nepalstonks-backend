package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iminfinity/nepalstonks/models"
	"github.com/iminfinity/nepalstonks/utils"
	"gopkg.in/mgo.v2/bson"
)

// GetStockList func
func GetStockList(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	var stockList models.StockList

	err = stocksListCollection.FindOne(ctx, bson.M{"bodge": "giveData"}).Decode(&stockList)
	if err != nil {
		http.Error(rw, "Failed getting stock list", http.StatusInternalServerError)
		fmt.Println("Failed getting stock list")
		return
	}
	json.NewEncoder(rw).Encode(&stockList.List)
	fmt.Println("Getting list of stock successfull")
}

// AddNewStock func
func AddNewStock(newStock string) bool {

	var stockList models.StockList

	err = stocksListCollection.FindOne(ctx, bson.M{"bodge": "giveData"}).Decode(&stockList)
	if err != nil {
		fmt.Println("Failed updating stock list")
		return true
	}

	list := stockList.List
	list = append(list, newStock)
	stockList.List = list

	_, err = stocksListCollection.UpdateOne(ctx, bson.M{"bodge": "giveData"}, bson.M{"$set": stockList})
	if err != nil {
		fmt.Println("Failed updating stock list")
		return true
	}

	return false
}

// IfNewThenAdd func
func IfNewThenAdd(stock string) {
	var stockList models.StockList

	err = stocksListCollection.FindOne(ctx, bson.M{"bodge": "giveData"}).Decode(&stockList)
	if err != nil {
		fmt.Println("Failed updating stock list")
	}

	list := stockList.List
	if check := utils.Contains(list, stock); check == false {
		list = append(list, stock)
	}

	stockList.List = list

	_, err = stocksListCollection.UpdateOne(ctx, bson.M{"bodge": "giveData"}, bson.M{"$set": stockList})
	if err != nil {
		fmt.Println("Failed updating stock list")
	}
}
