/*
Package resp
@Author: MoZhu
@File: errors
@Software: GoLand
*/
package resp

import (
	"fmt"
	"strings"
)

func NewError(code int, args ...interface{}) error {
	return &PredefinedError{
		code:       code,
		args:       args,
		templateZh: errorMap[code]["zh"],
		templateEn: errorMap[code]["en"],
		language:   "zh",
	}
}

func NewErrorWithMsg(code int, zh string, en string) error {
	return &PredefinedError{
		code:       code,
		templateZh: zh,
		templateEn: en,
		language:   "zh",
	}
}

type PredefinedError struct {
	code       int           // 错误码
	args       []interface{} // 错误参数
	templateZh string        // 中文的提示信息
	templateEn string        // 英文的提示信息
	language   string        // 请求的语言
}

func (e *PredefinedError) Error() string {
	if strings.Contains(strings.ToLower(e.language), "en") {
		return fmt.Sprintf(e.templateEn, e.args)
	}
	return fmt.Sprintf(e.templateZh, e.args)
}

func (e *PredefinedError) SetLanguage(language string) error {
	e.language = language
	return e
}

func (e *PredefinedError) Code() int {
	return e.code
}

// IsErrorEqual 比较 error 是否是某种 error 类型
func IsErrorEqual(err error, errCode int) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(*PredefinedError); !ok {
		return err.Error() == NewError(errCode).Error()
	}
	return err.(*PredefinedError).Code() == errCode
}
