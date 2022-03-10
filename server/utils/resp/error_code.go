/*
Package resp
@Author: MoZhu
@File: error_code
@Software: GoLand
*/
package resp

const (
	Success          = 0
	DatabaseError    = 2002
	InvalidParameter = 40000
	Unauthorized     = 40001
	RecordNotFound   = 40002
	ServerError      = 50000
)

const (
	ThirdPartyRequestTimeout       = 184001 // 请求超时
	ThirdPartyStatusCodeInvalid    = 184002 // httpclient 状态码不是 200
	ThirdPartyResponseInvalid      = 184003 // 返回内容 check 失败
	ThirdPartyRequestMethodInvalid = 184004 // httpclient 方法非法
	ThirdPartyReachRateLimited     = 184005 // 达到开放平台的频率限制
)
