package api

import (
	"github.com/gin-gonic/gin"
)

func Api() {
	router := gin.Default()

	router.POST("/auth/signup", Auth_Signup)
	router.POST("/auth/login", Auth_Login)
	router.GET("/oshimen", Get_Oshimen)
	router.POST("/oshimen", Post_Oshimen)
	router.DELETE("/oshimen/:id", Delete_Oshimen)
	router.GET("/oshimen/:id", Get_OshimenInfo)
	router.GET("/goods", Get_Goods)
	router.POST("/purchases", Post_Purchases)
	router.GET("/purchases", Get_Purchases)
	router.PUT("/purchases/:id", Put_Purchases)
	router.POST("/statistics", Post_Statistics)
	router.GET("/statistics", Get_Statistics)

	router.Run(":2020")
}
