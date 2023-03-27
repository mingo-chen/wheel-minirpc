package server

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/mingo-chen/wheel-minirpc/config"
	"github.com/mingo-chen/wheel-minirpc/logger"
	"github.com/mingo-chen/wheel-minirpc/transport"
)

// Startup 服务启动
func Startup(ctx context.Context, conf config.AppConfig) {
	// 启动Plugins

	// 启动Service
	for _, svr := range conf.Server.Service {
		go StartService(ctx, svr)
	}

	// 监听关闭信号

	// 优雅关闭服务
}

// StartService 启动服务
func StartService(ctx context.Context, service config.ServiceConfig) {
	defer func() {
		if e := recover(); e != nil {
			logger.ErrorContext(ctx, "Service panic: %+v, stack:%+v", e, string(debug.Stack()))
		}
	}()

	if service.Network == "tcp" {
		if err := transport.TcpServer(ctx, service.Port); err != nil {
			panic(fmt.Errorf("tcp serve err:%+v", err))
		}
	}
}
