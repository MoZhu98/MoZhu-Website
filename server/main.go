/*
@Author: MoZhu
@File: main
@Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mozhu98/website/server/global"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	global.Init()
}
