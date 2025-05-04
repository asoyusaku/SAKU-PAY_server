package model

type Goods struct {
	Name  string `json:"name" gorm:"primaryKey"`
	Price string `json:"price"`
	Image string `json:"image"`
}
