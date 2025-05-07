package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Api() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/auth/signup", Auth_Signup)
	// router.POST("/user", Add_User)              //complete
	router.GET("/user/:id", Get_User)           //complete
	router.GET("/oshimen/:id", Get_Oshimen)     //complete
	router.POST("/oshimen", Post_Oshimen)       //complete
	router.DELETE("/oshimen", Delete_Oshimen)   //complete
	router.GET("/goods", Get_AllGoods)          //complete
	router.GET("/purchases/:id", Get_Purchases) //complete
	router.POST("/purchases", Put_Purchases)    //complete
	router.GET("/cost/:id", Get_Total_Cost)     //complete

	router.Run(":9090")
}
