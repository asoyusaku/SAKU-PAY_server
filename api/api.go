package api

import (
	"github.com/gin-gonic/gin"
)

func Api() {
	router := gin.Default()

	router.POST("/auth/signup", Auth_Signup)  //complete
	router.POST("/user", Add_User)            //complete
	router.GET("/user/:id", Get_User)         //complete
	router.GET("/oshimen/:id", Get_Oshimen)   //complete
	router.POST("/oshimen", Post_Oshimen)     //complete
	router.DELETE("/oshimen", Delete_Oshimen) //complete
	router.GET("/goods", Get_AllGoods)        //complete
	router.POST("/goods", Post_Goods)         //complete
	router.GET("/goods/:id", Get_Goods)
	router.POST("/purchases", Post_Purchases)
	router.GET("/purchases", Get_Purchases)
	router.PUT("/purchases/", Put_Purchases)
	router.POST("/statistics", Post_Statistics)
	router.GET("/statistics", Get_Statistics)

	router.Run(":9090")
}
