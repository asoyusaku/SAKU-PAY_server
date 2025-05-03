package model

import (
	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	Name   string `json:"name"`
	Price  string `json:"price"`
	Image  string `json:"image"`
	UserID uint   `json:"user_id"`
}
