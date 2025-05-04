package api

import (
	"SAKU-PAY/database"
	"SAKU-PAY/model"
	"SAKU-PAY/scraping"

	"github.com/gin-gonic/gin"
)

// 新規登録
func Auth_Signup(c *gin.Context) { //complete
	var token model.IdToken

	if err := c.BindJSON(&token); err != nil {
		return
	}
	if err := database.AddIdToken(token); err != nil {
		c.JSON(500, gin.H{"error": "failed to add id token"})
		return
	}
}

func Add_User(c *gin.Context) { //complete
	var response model.Response

	if err := c.BindJSON(&response); err != nil {
		c.JSON(400, gin.H{"error": "failed to bind response"})
		return
	}

	if err := database.AddUser(response); err != nil {
		c.JSON(500, gin.H{"error": "failed to add user"})
		return
	}

	c.JSON(200, gin.H{"message": "user added successfully"})
}

func Get_User(c *gin.Context) { //complete
	id := c.Param("id")

	if user, err := database.GetUser(id); err != nil {
		c.JSON(500, gin.H{"error": "failed to get user"})
		return
	} else {
		c.JSON(200, gin.H{"user": user})
		return
	}
}

// 推しメン一覧取得
func Get_Oshimen(c *gin.Context) { //complete
	id := c.Param("id")

	if oshimen, err := database.GetOshimen(id); err != nil {
		c.JSON(500, gin.H{"error": "failed to get oshimen"})
		return
	} else {
		c.JSON(200, gin.H{"oshimen": oshimen})
		return
	}

}

// 推しメン追加
func Post_Oshimen(c *gin.Context) { //complete
	var request model.Request

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "failed to bind request"})
		return
	}
	if err := database.AddOshimen(request.UserId, request.Oshimen); err != nil {
		c.JSON(500, gin.H{"error": "failed to add oshimen"})
		return
	}
	c.JSON(200, gin.H{"message": "oshimen added successfully"})
}

// 推しメン削除
func Delete_Oshimen(c *gin.Context) { //complete
	var request model.Request

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "failed to bind request"})
		return
	}
	if err := database.DeleteOshimen(request.UserId, request.Oshimen); err != nil {
		c.JSON(500, gin.H{"error": "failed to delete oshimen"})
		return
	}
	c.JSON(200, gin.H{"message": "oshimen deleted successfully"})
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
