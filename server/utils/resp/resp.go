/*
Package resp
@Author: MoZhu
@File: resp
@Software: GoLand
*/
package resp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendData(c *gin.Context, data interface{}) {
	resp := &Response{}
	resp.Code = Success
	resp.Data = data
	c.JSON(http.StatusOK, resp)
}

func SendSuccess(c *gin.Context) {
	SendData(c, map[string]interface{}{"result": "success"})
}

func SendMsg(c *gin.Context, code int, msg string) {
	resp := &Response{}
	resp.Code = code
	resp.Message = msg
	c.JSON(http.StatusOK, resp)
}

func SendError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	var realError *PredefinedError
	errors.As(err, &realError)
	resp := &Response{}
	resp.Code = realError.Code()
	resp.Message = realError.SetLanguage(c.Query("language")).Error()
	c.JSON(http.StatusOK, resp)
}

type Response struct {
	Code    int         `json:"code"`    // 错误码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}
