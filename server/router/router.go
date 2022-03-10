/*
Package router
@Author: MoZhu
@File: router
@Software: GoLand
*/
package router

import (
	"github.com/mozhu98/website/server/handler/login"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/mozhu98/website/server/constant"
	"github.com/mozhu98/website/server/utils/logs"
)

type Server struct {
	engine   *gin.Engine
	loginAPI *login.API
}

type injector struct {
	dig.In
	Engine   *gin.Engine
	LoginAPI *login.API
}

func NewServer(injector injector) *Server {
	return &Server{
		engine:   injector.Engine,
		loginAPI: injector.LoginAPI,
	}
}

// InitServer 项目依赖注入
func InitServer(s *Server) {
	logs.Info(constant.ModuleInit, "Init server invoke %v", s != nil)
}

func (s *Server) Run(addr ...string) error {
	s.InitRouter()
	return s.engine.Run(addr...)
}

func (s *Server) InitRouter() *gin.Engine {
	router := s.engine
	router.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})
	router.Use(gin.RecoveryWithWriter(logs.PanicLog.Writer()))

	// 服务启动验证 ping
	router.GET("ping", func(context *gin.Context) {
		context.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<h1>PONG</h1>"))
	})
	s.Register()
	return router
}

func (s *Server) Register() {
	// 静态资源托管
	s.initStaticFile(s.engine.Group(""))
	group := s.engine.Group("api")

	s.initUnAuthGroup(group)
	s.initSSOGroup(group)
}

func (s *Server) initStaticFile(group *gin.RouterGroup) {

}

// 不需要认证的接口
func (s *Server) initUnAuthGroup(group *gin.RouterGroup) {
	{
		group.GET("login/:third_party", s.loginAPI.GetThirdPartyLoginInfo)         // 获取第三方 oauth 登录链接
		group.GET("login/callback/:third_party", s.loginAPI.ThirdPartyLoginByCode) // 消费第三方 oauth 登录回调 code
	}
}

// SSO IDP 相关接口实现
func (s *Server) initSSOGroup(group *gin.RouterGroup) {
	{
		casAuthGroup := group.Group("cas")
		casAuthGroup.Use()
		casAuthGroup.GET("login")
	}
	{
		casNoAuthGroup := group.Group("cas")
		casNoAuthGroup.GET("validate")
		casNoAuthGroup.GET("serviceValidate")
		casNoAuthGroup.GET("proxyValidate")
		casNoAuthGroup.GET("logout")
	}

	{
		oidcAuthGroup := group.Group("oidc")
		oidcAuthGroup.Use()
		oidcAuthGroup.GET("login")
	}
	{
		oidcNoAuthGroup := group.Group("oidc")
		oidcNoAuthGroup.POST("token")
		oidcNoAuthGroup.GET("userinfo")
		oidcNoAuthGroup.GET("logout")
	}

	{
		samlAuthGroup := group.Group("idp")
		samlAuthGroup.Use()
		samlAuthGroup.POST("sso")
		samlAuthGroup.GET("sso")
		samlAuthGroup.GET("entity/:sid")
	}
	{
		samlNoAuthGroup := group.Group("idp")
		samlNoAuthGroup.GET("metadata")
	}
}
