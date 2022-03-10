/*
Package third_party
@Author: MoZhu
@File: service
@Software: GoLand
*/
package third_party

import (
	"fmt"

	"github.com/mozhu98/website/server/dao"
	"github.com/mozhu98/website/server/model/req"
)

type Client interface {
	InitClientBySetting(settings map[string]string) error                 // 初始化
	GetLoginInfo() (map[string]string, error)                             // 生成三方登录信息
	GetUserByCode(r req.ThirdPartyCodeRequest) (*dao.User, string, error) // 获取登录用户信息
	Close()                                                               // 关闭
}

func TransThirdPartyAliasToEnum(tp string) ThirdPartyEnum {
	for k, v := range AliasMap {
		if tp == v {
			return k
		}
	}
	return UnknownType
}

func GetSettingKeyForSplicePrefix(alias ThirdPartyEnum, str string) string {
	return fmt.Sprintf("%v_%v_%v", SettingPrefix, AliasMap[alias], str)
}
