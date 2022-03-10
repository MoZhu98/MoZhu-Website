/*
Package config
@Author: MoZhu
@File: config
@Software: GoLand
*/
package config

import (
	"github.com/mozhu98/website/server/utils/env"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	Debug    bool   `yaml:"debug"`
	BindAddr string `yaml:"bind_addr"`
	Log      struct {
		EventLog  string `yaml:"event_log"`
		AccessLog string `yaml:"access_log"`
		PanicLog  string `yaml:"panic_log"`
	} `yaml:"log"`
}

var Config *config

func getConfigName() string {
	var configFile string
	environ := env.TellEnv()
	switch environ {
	case env.Online:
		configFile = "conf/online.yaml"
	case env.Release:
		configFile = "conf/release.yaml"
	case env.Dev:
		configFile = "conf/dev.yaml"
	case env.Loc:
		configFile = "conf/loc.yaml"

	}
	return configFile
}

func Init() {
	configFile := getConfigName()
	// 读取配置文件
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("read config file failed: %v", err)
	}
	err = yaml.Unmarshal(configData, &Config)
	if err != nil {
		log.Fatalf("unmarshal config file failed: %v", err)
	}
}
