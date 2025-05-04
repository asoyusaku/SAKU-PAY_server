package model

type IdToken struct {
	IdToken string `json:"id_token" gorm:"id_token"`
}

type Response struct {
	Iss     string `json:"iss"`
	Sub     string `json:"sub"`
	Aud     string `json:"aud"`
	Exp     int64  `json:"exp"`
	Iat     int64  `json:"iat"`
	Nonce   string `json:"nonce"`
	Amr     string `json:"amr"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Email   string `json:"email"`
}
