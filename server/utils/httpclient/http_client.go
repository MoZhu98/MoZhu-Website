/*
Package httpclient
@Author: MoZhu
@File: http_CLIENT
@Software: GoLand
*/
package httpclient

import (
	"errors"
	"math"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
	"go.uber.org/ratelimit"

	"github.com/mozhu98/website/server/constant"
	"github.com/mozhu98/website/server/utils/logs"
	"github.com/mozhu98/website/server/utils/resp"
)

const (
	RequestTimeout              = 20 * time.Second
	defaultTokenRefreshInterval = 20 * time.Second
	advanceTokenRefreshDuration = 20 * time.Second
)

var ErrRate = errors.New("open platform is limited")

type HTTPRequestOption struct {
	TargetURL   string
	Method      string
	Header      map[string]string
	Params      interface{}
	Body        interface{}
	ContentType string // 可选为 form、json
	retryTimes  int    // 本次重试次数
	retryErr    error  // 上次重试错误
	MaxRetry    int    // 最大重试次数
	AvoidRate   bool   // 是否绕过频率限制
}

func (opt *HTTPRequestOption) SetError(err error) {
	opt.retryTimes++
	opt.retryErr = err
}

type HTTPResponse interface {
	CheckResponse() error
}

type Client interface {
	Start()        // 启动客户端
	Close()        // 关闭客户端
	RefreshToken() // 定期刷新Token
	AddRater(limit int, duration time.Duration)
	Send(option *HTTPRequestOption, ptrOfStruct HTTPResponse) error
}

type DefaultClient struct {
	rater               []ratelimit.Limiter
	appendToken         func(*gorequest.SuperAgent, string) *gorequest.SuperAgent
	tokenGenerator      func() (string, time.Duration)
	token               string
	tokenRefreshCloseCh chan bool // 关闭 token 刷新的 channel
}

func (c *DefaultClient) Start() {
	if c.tokenGenerator != nil {
		c.RefreshToken()
	}
}

func (c *DefaultClient) Close() {
	if c.tokenRefreshCloseCh != nil {
		c.tokenRefreshCloseCh <- true
		close(c.tokenRefreshCloseCh)
	}
}

func (c *DefaultClient) RefreshToken() {
	next := c.getToken()
	go func() {
		for {
			next = c.getToken()
			t := time.NewTimer(next)
			select {
			case <-c.tokenRefreshCloseCh:
				return
			case <-t.C:
			}
		}
	}()
}

func (c *DefaultClient) AddRater(limit int, duration time.Duration) {
	c.rater = append(c.rater, ratelimit.New(limit, ratelimit.Per(duration)))
}

func (c *DefaultClient) Send(option *HTTPRequestOption, responseOfStruct HTTPResponse) error {
	// 校验是否达到频率上限
	if (option.MaxRetry == 0 && option.retryTimes > 5) || (option.MaxRetry != 0 && option.retryTimes > option.MaxRetry) {
		return option.retryErr
	} else if option.retryTimes != 0 {
		time.Sleep(time.Duration(math.Pow(2, math.Max(float64(option.retryTimes), 2))) * time.Second) // 递归
	}
	if !option.AvoidRate {
		for _, rater := range c.rater {
			rater.Take()
		}
	}
	start := time.Now()
	r := gorequest.New().Timeout(RequestTimeout)
	switch option.Method {
	case http.MethodGet:
		r = r.Get(option.TargetURL)
	case http.MethodPost:
		r = r.Post(option.TargetURL)
	default:
		return resp.NewError(resp.ThirdPartyRequestMethodInvalid, option.Method)
	}
	// 设置 Header
	if option.Header != nil {
		for k, v := range option.Header {
			r = r.AppendHeader(k, v)
		}
	}
	// 设置默认 Content-type
	if r.Header.Get("Content-Type") == "" {
		r = r.AppendHeader("Content-Type", "application/json")
	}
	// 设置 AccessToken
	if c.appendToken != nil {
		r = c.appendToken(r, c.token)
	}
	// 设置 Params
	if option.Params != nil {
		r = r.Query(option.Params)
	}
	// 设置 Body
	if option.Body != nil {
		r = r.SendStruct(option.Body)
	}

	var body []byte
	var response gorequest.Response
	var errs []error
	if responseOfStruct == nil {
		response, body, errs = r.EndBytes()
	} else {
		response, body, errs = r.EndStruct(responseOfStruct)
	}
	if errs != nil {
		for _, v := range errs {
			if v != nil {
				if v.Error() == "Client.Timeout exceeded while awaiting headers" {
					option.SetError(resp.NewError(resp.ThirdPartyRequestTimeout))
					return c.Send(option, responseOfStruct)
				}
				logs.Error(constant.ModuleHttp, "Request %v failed, errs is %v", option.TargetURL, v)
				return v
			}
		}
	}
	if response.StatusCode != http.StatusOK {
		logs.Error(constant.ModuleHttp, "Request %v status code is not 200, Body %v but %v", option.TargetURL, string(body), response.StatusCode)
		return resp.NewError(resp.ThirdPartyStatusCodeInvalid, response.StatusCode)
	}
	if responseOfStruct != nil {
		if err := responseOfStruct.CheckResponse(); err != nil {
			if err == ErrRate {
				option.SetError(resp.NewError(resp.ThirdPartyReachRateLimited, -1))
				return c.Send(option, responseOfStruct)
			}
			logs.Error(constant.ModuleHttp, "Request %v, Response Body %v is not valid, err is %v", option.TargetURL, string(body), err)
			return err
		}
	}
	logs.Debug(constant.ModuleHttp, "Request %v, Response Body %v, time %v", option.TargetURL, string(body), time.Now().Sub(start).String())
	return nil
}

func (c *DefaultClient) getToken() time.Duration {
	next := defaultTokenRefreshInterval
	token, expire := c.tokenGenerator()
	if token != "" {
		c.token = token
	}
	if expire > advanceTokenRefreshDuration {
		next = expire - advanceTokenRefreshDuration
	}
	return next
}
