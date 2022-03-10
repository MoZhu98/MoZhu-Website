/*
Package login
@Author: MoZhu
@File: login
@Software: GoLand
*/
package login

import (
	"github.com/mozhu98/website/server/dao"
	"github.com/mozhu98/website/server/model/req"
	"github.com/mozhu98/website/server/service/third_party"
	"go.uber.org/dig"
)

type Service interface {
	GetThirdPartyLoginInfo(tp third_party.ThirdPartyEnum) (map[string]string, error)
	GetUserByCode(tp third_party.ThirdPartyEnum, r req.ThirdPartyCodeRequest) (*dao.User, string, error)
}

type serviceImpl struct {
	settingDAO dao.SettingDAOIF
}

type injector struct {
	dig.In
	SettingDAO dao.SettingDAOIF
}

func NewService(injector injector) Service {
	return &serviceImpl{
		settingDAO: injector.SettingDAO,
	}
}
