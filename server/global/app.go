/*
Package global
@Author: MoZhu
@File: app
@Software: GoLand
*/
package global

import (
	"github.com/gin-gonic/gin"

	"github.com/mozhu98/website/server/config"
	"github.com/mozhu98/website/server/router"
)

// App 服务入口，包含 httpclient 接口，定时任务
type App struct {
	httpService *router.Server
}

func NewApp(httpService *router.Server) *App {
	return &App{
		httpService: httpService,
	}
}

func NewGinEngine() *gin.Engine {
	engine := gin.New()
	gin.SetMode(gin.ReleaseMode)
	return engine
}

// StartApp 启动服务
func StartApp(app *App) error {
	return app.httpService.Run(config.Config.BindAddr)
}
