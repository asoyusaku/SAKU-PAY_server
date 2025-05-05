package model

type Purchase struct {
	UserId    string `json:"user_id" gorm:"primaryKey"`
	GoodsName string `json:"goods_name" gorm:"primaryKey"`
	Quantity  int    `json:"quantity"`
}
