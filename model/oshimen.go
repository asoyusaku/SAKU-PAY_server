package model

import (
	"gorm.io/gorm"
)

type Oshimen struct {
	gorm.Model
	Name    string `json:"name"`
	Picture string `json:"picture"`
	UserID  uint   `json:"user_id"`
}

type Member struct {
	gorm.Model
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
