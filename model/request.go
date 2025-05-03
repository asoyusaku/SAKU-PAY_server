package model

type Request struct {
	UserId  string `json:"user_id"`
	Oshimen Member `json:"oshimen"`
}
