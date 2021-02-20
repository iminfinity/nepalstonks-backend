package models

// StockDetails struct
type StockDetails struct {
	Date         []string `json:"date,omitempty" bson:"date,omitempty"`
	MaxPrice     []string `json:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	MinPrice     []string `json:"minPrice,omitempty" bson:"minPrice,omitempty"`
	ClosingPrice []string `json:"closingPrice,omitempty" bson:"closingPrice,omitempty"`
}

// StockDataList struct
type StockDataList struct {
	StockName string       `json:"stockName,omitempty" bson:"stockName,omitempty"`
	StockData StockDetails `json:"stockData,omitempty" bson:"stockData,omitempty"`
}

// StockList struct
type StockList struct {
	List  []string `json:"list,omitempty" bson:"list,omitempty"`
	Bodge string   `json:"bodge,omitempty" bson:"bodge,omitempty"`
}

// ResponsedDataFromHerocu struct
type ResponsedDataFromHerocu struct {
	StockSymbol  string `json:"Symbol,omitempty"`
	ClosingPrice string `json:"Close,omitempty"`
	MinPrice     string `json:"Low,omitempty"`
	MaxPrice     string `json:"High,omitempty"`
}
