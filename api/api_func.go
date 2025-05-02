package api

import (
	database "SAKU-PAY/db"
	"SAKU-PAY/model"
	"SAKU-PAY/scraping"

	"github.com/gin-gonic/gin"
)

// 新規登録
func Auth_Signup(c *gin.Context) {
	var token model.IdToken

	if err := c.BindJSON(&token); err != nil {
		return
	}
	if err := database.AddIdToken(token); err != nil {
		c.JSON(500, gin.H{"error": "failed to add id token"})
		return
	}
}

// 推しメン一覧取得
func Get_Oshimen(c *gin.Context) {
}

// 推しメン追加
func Post_Oshimen(c *gin.Context) {
}

// 推しメン削除
func Delete_Oshimen(c *gin.Context) {
}

// 推しメン情報取得
func Get_OshimenInfo(c *gin.Context) {
}

// グッズ一覧取得
func Get_Goods(c *gin.Context) {
	scraping.Scrape_Goods()
}

// 購入記録追加
func Post_Purchases(c *gin.Context) {
}

// 購入記録取得
func Get_Purchases(c *gin.Context) {
}

// 購入記録更新
func Put_Purchases(c *gin.Context) {
}

// 購入集計記録
func Post_Statistics(c *gin.Context) {
}

// 購入集計取得
func Get_Statistics(c *gin.Context) {
}
