package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/",func(c *gin.Context){
		//map[string]interface{}
		// map[string]any{}
		c.JSON(200,gin.H{
			"message" : "Todo API is running!",
			"status" : "success",
		})
	})

	// start the server
	router.Run(":3000")
}
