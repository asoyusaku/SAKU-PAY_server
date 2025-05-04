package model

type User struct {
	ID      string   `json:"id" gorm:"primaryKey"` // IDトークンのuserid
	Name    string   `json:"name" gorm:"size:255"`
	Picture string   `json:"picture"`
	Email   string   `json:"email"`
	Oshimen []Member `json:"oshimen" gorm:"many2many:user_members;"`
	Goods   []Goods  `json:"goods" gorm:"many2many:user_goods;"`
}
