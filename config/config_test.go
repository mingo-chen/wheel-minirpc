package config

import (
	"testing"
)

func Test_load(t *testing.T) {
	LoadAppConf("./rpc.yaml")
	t.Logf("app config:%+v", App)

	for k1, c := range App.Plugins {
		t.Logf("k1:%+v, %+v", k1, c)
	}

	type logger struct {
		Level    string `yaml:"level"`
		Path     string `yaml:"path"`
		RollType string `yaml:"roll_type"`
		MaxSize  int    `yaml:"max_size"`
		Compress bool   `yaml:"compress"`
	}

	var log logger
	LoadPluginConf("logger", &log)
	// }
	t.Logf("--> %+v", log)
}
