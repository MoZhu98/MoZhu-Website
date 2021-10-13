/*
@Time : 2021/10/13 23:43
@Author : MoZhu
*/
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})
	router.Run(":8888")
}
