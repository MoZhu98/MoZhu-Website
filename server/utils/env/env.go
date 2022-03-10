/*
Package env
@Author: MoZhu
@File: env
@Software: GoLand
*/
package env

import (
	"os"
	"strings"
)

const (
	Env     = "MOZHU_ENV"
	Online  = "online"  // 线上环境
	Release = "release" // 测试环境
	Dev     = "dev"     // 联调环境
	Loc     = "loc"     // 本地
)

func TellEnv() string {
	environ := strings.ToLower(os.Getenv(Env))
	switch environ {
	case Online, Release, Dev:
		return environ
	default:
		return Loc
	}
}

func IsOnline() bool {
	return TellEnv() == Online
}

func IsRelease() bool {
	return TellEnv() == Release
}

func IsDev() bool {
	return TellEnv() == Dev
}

func IsLoc() bool {
	return TellEnv() == Loc
}
