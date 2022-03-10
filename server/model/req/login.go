/*
Package req
@Author: MoZhu
@File: login
@Software: GoLand
*/
package req

type ThirdPartyCodeRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}
