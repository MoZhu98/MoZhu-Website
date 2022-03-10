/*
Package global
@Author: MoZhu
@File: init
@Software: GoLand
*/
package global

import (
	"go.uber.org/dig"
	"log"

	"github.com/mozhu98/website/server/config"
	"github.com/mozhu98/website/server/router"
	"github.com/mozhu98/website/server/utils/env"
	"github.com/mozhu98/website/server/utils/logs"
)

func Init() {
	config.Init() // 初始化配置
	initLogger()  // 初始化日志系统
	c := dig.New()
	if err := Inject(c); err != nil {
		log.Fatalf("Ingect failed: %v", err)
	}
	if err := Invoke(c); err != nil {
		log.Fatalf("Invoke failed: %v", err)
	}
	// 确保依赖注入已生效
	if err := c.Invoke(router.InitServer); err != nil {
		log.Fatalf("Invoke init server failed: %v", err)
	}
	if err := c.Invoke(StartApp); err != nil {
		log.Fatalf("Start app failed: %v", err)
	}
}

func initLogger() {
	logConfig := logs.Config{
		LogPath:       config.Config.Log.EventLog,
		AccessLogPath: config.Config.Log.AccessLog,
		PanicLogPath:  config.Config.Log.PanicLog,
		LogLevel:      logs.LevelInfo,
		Stdout:        false,
	}
	if env.IsLoc() {
		logConfig.Stdout = true
		logConfig.LogLevel = logs.LevelDebug
	}
	if config.Config.Debug {
		logConfig.LogLevel = logs.LevelDebug
	}
	if err := logs.InitLogger(logConfig); err != nil {
		log.Fatalf("Init log failed: %v", err)
	}
}
