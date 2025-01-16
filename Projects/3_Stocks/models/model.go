package model


type Stock struct{
	Id int64 `json:"stockId"`
	Name string `json:"stockName"`
	Quantity int64 `json:"quantity"`
	Price int `json:"price"`
	Company string `json:"company"`
}