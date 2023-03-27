package plugins

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Plugin 插件接口，用于第三方采用合适的组件
type Plugin interface {
	// 启动插件
	Startup(cfg yaml.Node)
}

var pluginHub = map[string]Plugin{}

// GetPlugin 获取指定协议的帧编解码器
func GetPlugin(name string) Plugin {
	return pluginHub[name]
}

// RegistPlugin 注册帧编解码器
func RegistPlugin(name string, plugin Plugin) {
	pluginHub[name] = plugin
}

// Startup 启动所有插件
func Startup(cfgs map[string]yaml.Node) {
	for plg, cfg := range cfgs {
		plugin, ok := pluginHub[plg]
		if !ok {
			panic(fmt.Errorf("plugin[%s] not exist", plugin))
		}

		plugin.Startup(cfg)
	}
}
