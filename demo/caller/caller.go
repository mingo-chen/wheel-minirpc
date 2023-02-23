package main

import (
	"context"
	"fmt"
	"net"

	"github.com/mingo-chen/wheel-minirpc/demo"
	"github.com/mingo-chen/wheel-minirpc/demo/stub"
	"github.com/mingo-chen/wheel-minirpc/framework"
	"github.com/mingo-chen/wheel-minirpc/framework/ext"
)

func main() {
	framework.RegistCoder("minirpc", ext.MiniRpcReqCoder{}, ext.MiniRpcRspCoder{})
	framework.RegistFramer("minirpc", ext.MiniRpcFramer{})

	ctx := context.TODO()
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(fmt.Errorf("dail server err:%+v", err))
	}

	stub := stub.RpcDemoStub{
		Client: stub.RpcClient{
			Conn: conn,
		},
	}

	req := demo.HelloReq{Target: "zhangsan"}
	rsp, err := stub.Hello(ctx, &req)
	if err != nil {
		panic(fmt.Errorf("rpc[Hello] fail:%+v", err))
	}

	fmt.Printf("Get resp:%+v\n", rsp)
}
