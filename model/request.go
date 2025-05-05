package model

type Request_Oshimen struct {
	UserId  string `json:"user_id"`
	Oshimen Member `json:"oshimen"`
}

type Request_Purchase struct {
	UserId    string `json:"user_id"`
	GoodsName string `json:"goods_name"`
	Quantity  int    `json:"quantity"`
}
