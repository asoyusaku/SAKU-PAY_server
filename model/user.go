package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      string   `json:"id" gorm:"primaryKey"` // IDトークンのuserid
	Name    string   `json:"name" gorm:"size:255"`
	Picture string   `json:"picture"`
	Email   string   `json:"email"`
	Oshimen []Member `json:"oshimen" gorm:"many2many:user_members;"`
	Goods   []Goods  `json:"goods" gorm:"foreignKey:UserID"`
}
