package api

import (
	"SAKU-PAY/database"
	"SAKU-PAY/model"

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

// user情報を登録
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

// user情報を取得
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
	var request model.Request_Oshimen

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
	var request model.Request_Oshimen

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

// グッズ一覧取得
func Get_AllGoods(c *gin.Context) { //complete
	elements_goods, err := database.GetAllGoods()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get goods"})
		return
	}
	c.JSON(200, gin.H{"goods": elements_goods})
}

// 購入記録取得
func Get_Purchases(c *gin.Context) { //complete
	id := c.Param("id")
	if goods, err := database.GetGoods(id); err != nil {
		c.JSON(500, gin.H{"error": "failed to get goods"})
		return
	} else {
		c.JSON(200, gin.H{"goods": goods})
		return
	}
}

// 購入記録追加または更新
func Put_Purchases(c *gin.Context) { //complete
	var purchase model.Request_Purchase
	if err := c.BindJSON(&purchase); err != nil {
		c.JSON(400, gin.H{"error": "failed to bind purchase"})
		return
	}
	if err := database.UpdateGoods(purchase); err != nil {
		c.JSON(500, gin.H{"error": "failed to update goods"})
		return
	}
	c.JSON(200, gin.H{"message": "goods updated successfully"})
}

// 購入集計取得
func Get_Total_Cost(c *gin.Context) { //complete
	id := c.Param("id")
	if total_cost, err := database.Get_TotalCost(id); err != nil {
		c.JSON(500, gin.H{"error": "failed to get total cost"})
		return
	} else {
		c.JSON(200, gin.H{"total_cost": total_cost})
		return
	}
}
