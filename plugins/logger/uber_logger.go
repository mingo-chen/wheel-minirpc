package logger

import (
	"context"
	"fmt"

	"github.com/mingo-chen/wheel-minirpc/logger"
	"gopkg.in/yaml.v3"
)

type LoggerPlugin struct {
}

// load config by plugin name
func (l LoggerPlugin) Startup(cfg yaml.Node) {
	var log miniLog
	if err := cfg.Decode(&log.config); err != nil {
		panic(fmt.Errorf("decode plugin config err:%+v", err))
	}

	// 生成logger对象，注入到logger.Log中
	logger.Log = log
}

// 日志配置
type logConfig struct {
	Level    string `yaml:"level"`     // 日志级别
	Path     string `yaml:"path"`      // 输出地址
	RollType string `yaml:"roll_type"` // 日志滚动切割方式，按大小，按时间
	MaxSize  int    `yaml:"max_size"`  // 按大小切割时的上限
	Compress bool   `yaml:"compress"`  // 日志切割后是否要压缩
	Format   string `yaml:"format"`    // 日志格式
}

// miniLog 框架自带实现的日志器
type miniLog struct {
	config logConfig
}

// Debug Debug级别日志输出
func (d miniLog) Debug(format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}

// DebugContext Debug级别日志输出
func (d miniLog) DebugContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}

// Info Info级别日志输出
func (d miniLog) Info(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

// InfoContext Info级别日志输出
func (d miniLog) InfoContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

// Warn Warn级别日志输出
func (d miniLog) Warn(format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format+"\n", args...)
}

// WarnContext Warn级别日志输出
func (d miniLog) WarnContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format+"\n", args...)
}

// Error Error级别日志输出
func (d miniLog) Error(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}

// ErrorContext Error级别日志输出
func (d miniLog) ErrorContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}
