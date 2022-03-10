/*
Package third_party
@Author: MoZhu
@File: GitHub
@Software: GoLand
*/
package third_party

import (
	"github.com/mozhu98/website/server/dao"
	"github.com/mozhu98/website/server/model/req"
	httpUtil "github.com/mozhu98/website/server/utils/http"
	"github.com/mozhu98/website/server/utils/httpclient"
	"net/http"
)

type GitHubClient struct {
	NameAlias    ThirdPartyEnum
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AccessToken  string
	client       httpclient.Client
}

func (c *GitHubClient) GenSettingKey(key string) string {
	return GetSettingKeyForSplicePrefix(c.NameAlias, key)
}

func (c *GitHubClient) InitClientBySetting(settings map[string]string) error {
	if c.NameAlias != GitHub {
		panic("Client alias type error")
	}
	c.ClientID = settings[c.GenSettingKey("client_id")]
	c.ClientSecret = settings[c.GenSettingKey("client_secret")]
	c.RedirectURL = "http://127.0.0.1:8888" + CallbackPrefixURL + AliasMap[c.NameAlias]
	c.client = &httpclient.DefaultClient{}
	c.client.Start()
	return nil
}

func (c *GitHubClient) GetLoginInfo() (map[string]string, error) {
	result := make(map[string]string)
	result["alias"] = AliasMap[c.NameAlias]
	result["login_url"] = GitHubAuthorizeURL + "?" + httpUtil.Util.GetParamsSplicingEncode(map[string]string{
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
		"redirect_uri":  c.RedirectURL,
	})
	return result, nil
}

func (c *GitHubClient) GetUserByCode(r req.ThirdPartyCodeRequest) (*dao.User, string, error) {
	var responseToken GitHubAccessTokenResponse
	var err error
	err = c.client.Send(&httpclient.HTTPRequestOption{
		TargetURL: GitHubGetAccessTokenURL,
		Method:    http.MethodPost,
		Header: map[string]string{
			"Accept": "application/json",
		},
		Body: map[string]interface{}{
			"client_id":     c.ClientID,
			"client_secret": c.ClientSecret,
			"code":          r.Code,
		},
	}, &responseToken)
	if err != nil {
		return nil, "", err
	}
	var responseUser GitHubUserResponse
	err = c.client.Send(&httpclient.HTTPRequestOption{
		TargetURL: GitHubGetUserURL,
		Method:    http.MethodGet,
		Header: map[string]string{
			"Authorization": "token " + responseToken.AccessToken,
		},
	}, &responseUser)
	if err != nil {
		return nil, "", err
	}
	return GenerateUserByGitHub(responseUser), "", err
}

func (c *GitHubClient) Close() {
	c.client.Close()
}

type GitHubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func (r *GitHubAccessTokenResponse) CheckResponse() error {
	return nil
}

type GitHubUserResponse struct {
	Login     string `json:"login"`
	NodeID    string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

func (r *GitHubUserResponse) CheckResponse() error {
	return nil
}

func GenerateUserByGitHub(user GitHubUserResponse) *dao.User {
	return &dao.User{
		UserName:  user.Login,
		UserID:    user.NodeID,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
		Source:    int(GitHub),
	}
}
