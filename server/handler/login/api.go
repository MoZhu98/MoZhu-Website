/*
Package login
@Author: MoZhu
@File: api
@Software: GoLand
*/
package login

import (
	"github.com/mozhu98/website/server/service/login"
	"go.uber.org/dig"
)

type API struct {
	loginService login.Service
}

type injector struct {
	dig.In
	LoginService login.Service
}

func NewAPI(injector injector) *API {
	return &API{
		loginService: injector.LoginService,
	}
}
