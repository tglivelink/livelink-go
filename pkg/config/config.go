package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

/*
全局配置管理，
可以通过设置DefaultConfigLoader，自定义配置加载器
*/

type Config struct {
	Domain string `json:"domain" yaml:"domain"`   // 请求地址
	Appid  string `json:"appid" yaml:"appid"`     // 请求方标识
	SigKey string `json:"sig_key" yaml:"sig_key"` // 用于计算签名
	SecKey string `json:"sec_key" yaml:"sec_key"` // 用于计算用户加解密信息

	Signer     int `json:"signer" yaml:"signer"`         // 签名方式
	Coder      int `json:"coder" yaml:"coder"`           // 用户信息加解密方式
	Serializer int `json:"serializer" yaml:"serializer"` // 序列化方式
}

// GlobalConfig 获取全局配置
func GlobalConfig() *Config {
	return DefaultConfigLoader.Load()
}

var ConfigPath = "./livelink.yaml"

var DefaultConfigLoader ConfigLoader = &config{cfg: &Config{}, once: sync.Once{}}

type ConfigLoader interface {
	Load() *Config
}

type config struct {
	cfg  *Config
	once sync.Once
}

func (c *config) Load() *Config {
	c.once.Do(func() {
		bs, err := os.ReadFile(ConfigPath)
		if err != nil {
			panic(fmt.Sprintf("Load Config os.ReadFile err %v", err))
		}
		switch ext := filepath.Ext(ConfigPath); {
		case strings.Contains(ext, "yaml"):
			err = yaml.Unmarshal(bs, c.cfg)
		case strings.Contains(ext, "json"):
			err = json.Unmarshal(bs, c.cfg)
		default:
			err = fmt.Errorf("unsupport file type %v", ext)
		}
		if err != nil {
			panic(fmt.Sprintf("Load Config Unmarshal err %v", err))
		}
	})
	return c.cfg
}
