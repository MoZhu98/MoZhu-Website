/*
Package login
@Author: MoZhu
@File: third_party
@Software: GoLand
*/
package login

import (
	"github.com/gin-gonic/gin"
	"github.com/mozhu98/website/server/model/req"
	"github.com/mozhu98/website/server/service/third_party"
	"github.com/mozhu98/website/server/utils/resp"
	"net/http"
)

func (api *API) GetThirdPartyLoginInfo(c *gin.Context) {
	tp := c.Param("third_party")
	result, err := api.loginService.GetThirdPartyLoginInfo(third_party.TransThirdPartyAliasToEnum(tp))
	if err != nil {
		resp.SendError(c, err)
		return
	}
	//resp.SendData(c, result)
	c.Redirect(http.StatusFound, result["login_url"])
}

func (api *API) ThirdPartyLoginByCode(c *gin.Context) {
	tp := c.Param("third_party")
	r := req.ThirdPartyCodeRequest{}
	err := c.ShouldBindQuery(&r)
	if err != nil || tp == "" {
		resp.SendError(c, resp.NewError(resp.InvalidParameter))
		return
	}
	user, _, err := api.loginService.GetUserByCode(third_party.TransThirdPartyAliasToEnum(tp), r)
	if err != nil {
		return
	}
	resp.SendData(c, user)
}
