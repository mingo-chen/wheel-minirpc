package logger

import (
	"context"
	"fmt"
)

// 默认日志器实现，输出到控制台，用于单机调试
type defaultLogger struct{}

// Debug Debug级别日志输出
func (d defaultLogger) Debug(format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}

// DebugContext Debug级别日志输出
func (d defaultLogger) DebugContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}

// Info Info级别日志输出
func (d defaultLogger) Info(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

// InfoContext Info级别日志输出
func (d defaultLogger) InfoContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

// Warn Warn级别日志输出
func (d defaultLogger) Warn(format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format+"\n", args...)
}

// WarnContext Warn级别日志输出
func (d defaultLogger) WarnContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format+"\n", args...)
}

// Error Error级别日志输出
func (d defaultLogger) Error(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}

// ErrorContext Error级别日志输出
func (d defaultLogger) ErrorContext(ctx context.Context, format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}
