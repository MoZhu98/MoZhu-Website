/*
Package resp
@Author: MoZhu
@File: error_msg
@Software: GoLand
*/
package resp

var (
	errorMap = map[int]map[string]string{
		InvalidParameter: {
			"zh": "参数错误",
			"en": "Parameter is invalid.",
		},
		DatabaseError: {
			"zh": "数据库错误",
			"en": "DataBase Error",
		},
		Unauthorized: {
			"zh": "未授权的请求",
			"en": "Unauthorized request.",
		},
		RecordNotFound: {
			"zh": "对象不存在",
			"en": "Record not found.",
		},
		ServerError: {
			"zh": "服务器错误",
			"en": "Server error.",
		},
		ThirdPartyRequestTimeout: {
			"zh": "开放平台请求超时，请联系管理员或重试（184001）",
			"en": "Open platform request timeout, please content the enterprise administrator or retry (184001)",
		},
		ThirdPartyStatusCodeInvalid: {
			"zh": "开放平台请求错误，异常状态码 %v，请联系管理员（184002）",
			"en": "Open platform request error, the error code is %v, please content the enterprise administrator (184002)",
		},
		ThirdPartyResponseInvalid: {
			"zh": "开放平台请求错误，异常状态码 %v，异常信息 %v，请联系管理员（184003）",
			"en": "Open platform request error, the error code is %v, the error message is %v, please content the enterprise administrator (184003)",
		},
		ThirdPartyRequestMethodInvalid: {
			"zh": "非法的开放平台请求方法: %v",
			"en": "Invalid request method: %v",
		},
		ThirdPartyReachRateLimited: {
			"zh": "开放平台达到频率限制，错误码 %v",
			"en": "Open platform is limited, code is %v",
		},
	}
)
