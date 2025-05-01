package model

import (
	"gorm.io/gorm"
)

type Oshimen struct {
	gorm.Model
	Name       string `json:"name" gorm:"primaryKey"`
	Age        int    `json:"age"`
	Generation string `json:"generation"`
}

type OshimenList struct {
	Oshimen []Oshimen `json:"oshimen"`
}
