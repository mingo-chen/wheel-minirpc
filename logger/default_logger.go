package logger

import (
	"context"
	"fmt"
)

// 默认日志器实现，输出到控制台，用于单机调试
type stdLogger struct{}

// Debug Debug级别日志输出
func (d stdLogger) Debug(format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}

// DebugContext Debug级别日志输出
func (d stdLogger) DebugContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}

// Info Info级别日志输出
func (d stdLogger) Info(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

// InfoContext Info级别日志输出
func (d stdLogger) InfoContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

// Warn Warn级别日志输出
func (d stdLogger) Warn(format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format+"\n", args...)
}

// WarnContext Warn级别日志输出
func (d stdLogger) WarnContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format+"\n", args...)
}

// Error Error级别日志输出
func (d stdLogger) Error(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}

// ErrorContext Error级别日志输出
func (d stdLogger) ErrorContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}
