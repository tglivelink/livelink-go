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

type ServerConfig struct {
	Domain string `json:"domain" yaml:"domain"`
}

type ClientConfig struct {
	Appid  string `json:"appid" yaml:"appid"`
	SigKey string `json:"sig_key" yaml:"sig_key"`
	SecKey string `json:"sec_key" yaml:"sec_key"`
}

type Config struct {
	Server *ServerConfig `json:"server" yaml:"server"`
	Client *ClientConfig `json:"client" yaml:"client"`
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

/********************/

func SetGlobalConfig(cfg *Config) {
	c := &config{
		cfg:  cfg,
		once: sync.Once{},
	}
	c.once.Do(func() {})
	DefaultConfigLoader = c
}
