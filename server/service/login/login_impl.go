/*
Package login
@Author: MoZhu
@File: login_impl
@Software: GoLand
*/
package login

import (
	"github.com/mozhu98/website/server/constant"
	"github.com/mozhu98/website/server/dao"
	"github.com/mozhu98/website/server/model/req"
	"github.com/mozhu98/website/server/service/third_party"
	"github.com/mozhu98/website/server/utils/logs"
	"github.com/mozhu98/website/server/utils/resp"
)

func (s *serviceImpl) GetThirdPartyLoginInfo(tp third_party.ThirdPartyEnum) (map[string]string, error) {
	client, err := s.GenerateThirdPartyClient(tp)
	if err != nil {
		return nil, err
	}
	return client.GetLoginInfo()
}

func (s *serviceImpl) GetUserByCode(tp third_party.ThirdPartyEnum, r req.ThirdPartyCodeRequest) (*dao.User, string, error) {
	client, err := s.GenerateThirdPartyClient(tp)
	if err != nil {
		return nil, "", err
	}
	return client.GetUserByCode(r)
}

func (s *serviceImpl) GenerateThirdPartyClient(tp third_party.ThirdPartyEnum) (third_party.Client, error) {
	settings, err := s.settingDAO.GetSettings()
	if err != nil {
		logs.Error(constant.ModuleSetting, "Get settings failed: %v", err)
		return nil, resp.NewError(resp.DatabaseError)
	}
	var client third_party.Client
	switch tp {
	case third_party.GitHub:
		client = &third_party.GitHubClient{NameAlias: tp}
	default:

	}
	err = client.InitClientBySetting(settings)
	if err != nil {
		return nil, err
	}
	return client, nil
}
