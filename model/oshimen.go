package model

type Member struct {
	Name    string `json:"name" gorm:"primaryKey"`
	Picture string `json:"picture"`
}
