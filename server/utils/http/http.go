/*
Package http
@Author: MoZhu
@File: http
@Software: GoLand
*/
package http

import "net/url"

var Util Handle

type Handle struct{}

func (h *Handle) GetParamsSplicingEncode(param map[string]string) string {
	var params url.URL
	urlStr := params.Query()
	for k, v := range param {
		urlStr.Add(k, v)
	}
	return urlStr.Encode()
}
