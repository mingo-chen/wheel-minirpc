package logger

import (
	"context"
)

// 注入外部日志实现
var logger Logger = defaultLogger{}

// Logger接口
type Logger interface {
	Debug(format string, args ...interface{})
	DebugContext(ctx context.Context, format string, args ...interface{})

	Info(format string, args ...interface{})
	InfoContext(ctx context.Context, format string, args ...interface{})

	Warn(format string, args ...interface{})
	WarnContext(ctx context.Context, format string, args ...interface{})

	Error(format string, args ...interface{})
	ErrorContext(ctx context.Context, format string, args ...interface{})
}

// Debug Debug级别日志输出
func Debug(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

// DebugContext Debug级别日志输出
func DebugContext(ctx context.Context, format string, args ...interface{}) {
	logger.DebugContext(ctx, format, args...)
}

// Info Info级别日志输出
func Info(format string, args ...interface{}) {
	logger.Info(format, args...)
}

// InfoContext Info级别日志输出
func InfoContext(ctx context.Context, format string, args ...interface{}) {
	logger.InfoContext(ctx, format, args...)
}

// Warn Warn级别日志输出
func Warn(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

// WarnContext Warn级别日志输出
func WarnContext(ctx context.Context, format string, args ...interface{}) {
	logger.WarnContext(ctx, format, args...)
}

// Error Error级别日志输出
func Error(format string, args ...interface{}) {
	logger.Error(format, args...)
}

// ErrorContext Error级别日志输出
func ErrorContext(ctx context.Context, format string, args ...interface{}) {
	logger.ErrorContext(ctx, format, args...)
}
