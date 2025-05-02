package model

import (
	"gorm.io/gorm"
)

type Oshimen struct {
	gorm.Model
	Name    string `json:"name" gorm:"primaryKey"`
	Picture string `json:"picture"`
}

type OshimenList struct {
	Oshimen []Oshimen `json:"oshimen"`
}

type Member struct {
	gorm.Model
	Name    string `json:"name" gorm:"primaryKey"`
	Picture string `json:"picture"`
}
