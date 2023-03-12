package stub

import (
	"context"
	"net"

	"github.com/mingo-chen/wheel-minirpc/core"
	"github.com/mingo-chen/wheel-minirpc/demo"
	"github.com/mingo-chen/wheel-minirpc/transport"
	"google.golang.org/protobuf/proto"
)

// RegsiterImpl 后续通过code generate技术自动生成
func RegsiterImpl(ctx context.Context, impl rpcDemo) {
	transport.RegistRpc("Hello", func(ctx context.Context, req []byte) ([]byte, error) {
		var helleReq demo.HelloReq
		err := proto.Unmarshal(req, &helleReq)
		if err != nil {
			return nil, err
		}

		rsp, err := impl.Hello(ctx, &helleReq)
		if err != nil {
			return nil, err
		}
		return proto.Marshal(rsp)
	})
}

type rpcDemo interface {
	Hello(ctx context.Context, req *demo.HelloReq) (*demo.HelloRsp, error)
}

type RpcClient struct {
	Conn net.Conn
}

func (cli RpcClient) Invoke(ctx context.Context, req proto.Message, rsp proto.Message) (err error) {
	framer := core.GetFramer("minirpc")            // TODO: 读取配置
	reqCoder, rspCoder := core.GetCoder("minirpc") // TODO: 读取配置
	rpcCtx := core.RpcContext(ctx)

	// 生成reqId

	body, _ := proto.Marshal(req)
	reqData, err := reqCoder.Encode(rpcCtx, body)
	if err != nil {
		return err
	}

	// 往服务端发包
	if err := framer.WriteFrame(cli.Conn, reqData); err != nil {
		return err
	}

	if rspData, err := framer.ReadFrame(cli.Conn); err != nil {
		return err
	} else {
		if reply, err := rspCoder.Decode(rpcCtx, rspData); err != nil {
			return err
		} else {
			err = proto.Unmarshal(reply, rsp)
			return err
		}
	}
}

// 客户端代码
type RpcDemoStub struct {
	Client RpcClient
}

func (stub RpcDemoStub) Hello(ctx context.Context, req *demo.HelloReq) (*demo.HelloRsp, error) {
	// 设置上下文信息
	rpcCtx := core.NewDefaultRpcCtx()
	rpcCtx.WithName("Hello")
	ctx = core.WithCtx(ctx, rpcCtx)

	var rsp demo.HelloRsp
	if err := stub.Client.Invoke(ctx, req, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}
