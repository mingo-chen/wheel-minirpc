package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mingo-chen/wheel-minirpc/demo"
	"github.com/mingo-chen/wheel-minirpc/demo/stub"
	"github.com/mingo-chen/wheel-minirpc/framework"
	"github.com/mingo-chen/wheel-minirpc/framework/ext"
	"github.com/mingo-chen/wheel-minirpc/framework/transport"
)

func main() {
	framework.RegistCoder("minirpc", ext.MiniRpcReqCoder{}, ext.MiniRpcRspCoder{})
	framework.RegistFramer("minirpc", ext.MiniRpcFramer{})

	ctx := context.TODO()
	stub.RegsiterImpl(ctx, rpcDemoImpl{})

	if err := transport.TcpServer(ctx, 8080); err != nil {
		panic(fmt.Errorf("tcp serve err:%+v", err))
	}
}

type rpcDemoImpl struct{}

func (impl rpcDemoImpl) Hello(ctx context.Context, req *demo.HelloReq) (*demo.HelloRsp, error) {
	if len(req.Target) == 0 {
		return nil, errors.New("target must hava value")
	}

	return &demo.HelloRsp{Result: fmt.Sprintf("Hi, %s", req.Target)}, nil
}
