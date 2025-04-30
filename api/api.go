package api

import (
	"github.com/gin-gonic/gin"
)

// 新規登録
func Auth_Signup(c *gin.Context) {
}

// ログイン
func Auth_Login(c *gin.Context) {
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
