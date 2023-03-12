package transport

import (
	"context"
	"fmt"
	"net"

	"github.com/mingo-chen/wheel-minirpc/core"
)

func TcpServer(ctx context.Context, port int) error {
	lister, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return err
	}

	for {
		conn, err := lister.Accept()
		if err != nil {
			continue
		}

		go handConn(ctx, conn)
	}
}

func handConn(ctx context.Context, conn net.Conn) {
	framer := core.GetFramer("minirpc") // TODO: 读取配置
	req, err := framer.ReadFrame(conn)
	if err != nil {
		fmt.Printf("ReadFrame err:%+v \n", err)
	}

	// rpc handler
	rsp, err := handRpc(ctx, req)
	if err != nil {
		fmt.Printf("handRpc err:%+v \n", err)
	}

	err = framer.WriteFrame(conn, rsp)
	if err != nil {
		fmt.Printf("WriteFrame err:%+v \n", err)
	}
}

func handRpc(ctx context.Context, req []byte) (rsp []byte, err error) {
	reqCoder, rspCoder := core.GetCoder("minirpc") // TODO: 读取配置
	rpcCtx := core.RpcContext(ctx)
	request, err := reqCoder.Decode(rpcCtx, req)
	if err != nil {
		return nil, err
	}

	rpcName := rpcCtx.GetName()
	impl := getRpcImpl(rpcName) // get handler implements by rpcName
	if impl == nil {
		return nil, fmt.Errorf("rpc[%s] has no implement", rpcName)
	}

	ctx = core.WithCtx(ctx, rpcCtx)
	response, err := impl(ctx, request)
	if err != nil { // write to header, pass to client
		rpcCtx.WithExt("ret", 500)
		rpcCtx.WithExt("msg", err.Error())
	}

	rsp, err = rspCoder.Encode(rpcCtx, response)
	return
}

type rpcImpl func(ctx context.Context, req []byte) ([]byte, error)

func getRpcImpl(rpcName string) rpcImpl {
	return rpcHub[rpcName]
}

var rpcHub = map[string]rpcImpl{}

func RegistRpc(name string, impl rpcImpl) {
	rpcHub[name] = impl
}
