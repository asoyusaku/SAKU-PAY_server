package model

type Response_Goods struct {
	Goods    Goods `json:"goods"`
	Quantity int   `json:"quantity"`
}
