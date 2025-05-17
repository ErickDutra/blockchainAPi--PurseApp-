package controller

import "github.com/gin-gonic/gin"

func BlockChainController() {
	server := gin.Default()

	server.GET("/blockchain", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, Blockchain!",
		})

	})

	server.Run(":8080")
}