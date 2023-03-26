package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mingo-chen/wheel-minirpc/config"
	"github.com/mingo-chen/wheel-minirpc/demo"
	_ "github.com/mingo-chen/wheel-minirpc/demo/adapter"
	"github.com/mingo-chen/wheel-minirpc/demo/stub"
	"github.com/mingo-chen/wheel-minirpc/server"
)

func main() {
	ctx := context.TODO()
	if err := config.LoadAppConf("./server.yaml"); err != nil {
		panic(fmt.Errorf("load config err:%+v", err))
	}

	stub.RegsiterImpl(ctx, rpcDemoImpl{})

	server.Startup(ctx, config.App) // 阻塞执行
}

type rpcDemoImpl struct{}

func (impl rpcDemoImpl) Hello(ctx context.Context, req *demo.HelloReq) (*demo.HelloRsp, error) {
	if len(req.Target) == 0 {
		return nil, errors.New("target must hava value")
	}

	return &demo.HelloRsp{Result: fmt.Sprintf("Hi, %s", req.Target)}, nil
}
