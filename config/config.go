package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// App 应用配置对象
var App AppConfig

// 整个应用配置对象
type AppConfig struct {
	Global  GlobalConfig         `yaml:"global"`
	Server  ServerConfig         `yaml:"server"`
	Client  ClientConfig         `yaml:"client"`
	Plugins map[string]yaml.Node `yaml:"plugin"`
}

// 整个App共享的配置
// dev
// test
//
//	|- env: 调试
//	|- env: base
//	|- env: 集成
//
// prod
//
//	 |- env: pre
//		|- env: formal
//			|- set: gz
//			|- set: bj
type GlobalConfig struct {
	Namespace string `yaml:"namespace"` // 当前运行环境，dev/test/production
	EnvName   string `yaml:"env_name"`  // 当前的环境名，比如预发
	SetName   string `yaml:"set_name"`  //
}

// 对外提供服务配置声明
type ServiceConfig struct {
	Name     string `yaml:"name"`
	Port     int    `yaml:"port"`
	Network  string `yaml:"network"`
	Protocol string `yaml:"protocol"`
	Timeout  int    `yaml:"timeout"`
}

type ServerConfig struct {
	Name    string          `yaml:"name"`
	Filter  []string        `yaml:"filter"`
	Service []ServiceConfig `yaml:"service"`
}

type ApiConfig struct {
	Name   string `yaml:"name"`
	Timeut int    `yaml:"timeut"`
}

// 依赖的外部服务配置声明
type OutServiceConfig struct {
	Name     string    `yaml:"name"`
	Endpoint string    `yaml:"endpoint"`
	Protocol string    `yaml:"protocol"`
	Timeout  int       `yaml:"timeout"`
	API      ApiConfig `yaml:"api"`
}

type ClientConfig struct {
	Filter  []string         `yaml:"filter"`
	Timeout int              `yaml:"timeout"`
	Service OutServiceConfig `yaml:"service"`
}

// LoadAppConf 加载应用配置
func LoadAppConf(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("read yaml err:%+v", err)
		return err
	}

	if yaml.Unmarshal(content, &App); err != nil {
		fmt.Printf("Unmarshal err:%+v", err)
		return err
	}
	return nil
}

// LoadPluginConf 加载插件配置
func LoadPluginConf(name string, holder interface{}) error {
	cfg, ok := App.Plugins[name]
	if !ok {
		return fmt.Errorf("plugin config not exist")
	}

	return cfg.Decode(holder)
}
