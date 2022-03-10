/*
Package global
@Author: MoZhu
@File: inject
@Software: GoLand
*/
package global

import (
	"github.com/mozhu98/website/server/handler/login"
	"github.com/mozhu98/website/server/router"
	loginService "github.com/mozhu98/website/server/service/login"
	"github.com/mozhu98/website/server/service/session"
	"github.com/pkg/errors"
	"go.uber.org/dig"

	"github.com/mozhu98/website/server/dao"
)

// Inject 将 provider 在 dig Container 注册
func Inject(c *dig.Container) (err error) {
	// 注册 Provider
	err = c.Provide(NewApp)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "NewApp")
	}
	err = c.Provide(NewGinEngine)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "NewGinEngine")
	}
	err = c.Provide(router.NewServer)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "router.NewServer")
	}
	err = c.Provide(login.NewAPI)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "login.NewAPI")
	}
	err = c.Provide(loginService.NewService)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "login.NewService")
	}
	err = c.Provide(session.NewService)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "session.NewService")
	}
	err = c.Provide(dao.NewDBClient)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "dao.NewDBClient")
	}
	err = c.Provide(dao.NewSettingDAO)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "dao.NewSettingDAO")
	}
	err = c.Provide(dao.NewSessionDAO)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "dao.NewSessionDAO")
	}
	err = c.Provide(dao.NewTodoDAO)
	if err != nil {
		return errors.WithMessagef(err, "Add provider %s failed", "dao.NewTodoDAO")
	}
	return

}

// Invoke 将依赖注入，执行初始化操作
func Invoke(c *dig.Container) (err error) {
	return
}
