/*
Package third_party
@Author: MoZhu
@File: constant
@Software: GoLand
*/
package third_party

import (
	"strconv"
)

type ThirdPartyEnum int

func (e ThirdPartyEnum) S() string {
	return strconv.Itoa(int(e))
}

const (
	SettingPrefix     string = "third_party"
	CallbackPrefixURL string = "/api/login/callback/"
)

// 三方登录
const (
	UnknownType ThirdPartyEnum = iota
	GitHub
)

var AliasMap = map[ThirdPartyEnum]string{
	GitHub: "github",
}

const (
	GitHubAuthorizeURL      string = "https://github.com/login/oauth/authorize"
	GitHubGetAccessTokenURL string = "https://github.com/login/oauth/access_token"
	GitHubGetUserURL        string = "https://api.github.com/user"
)
