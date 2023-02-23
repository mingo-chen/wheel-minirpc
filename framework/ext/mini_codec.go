package ext

import (
	"fmt"
	"sync/atomic"

	"github.com/mingo-chen/wheel-minirpc/framework"
	miniproto "github.com/mingo-chen/wheel-minirpc/framework/ext/proto"
	"google.golang.org/protobuf/proto"
)

type MiniRpcReqCoder struct {
}

var reqId uint32

// 把请求上下文及业务请求包，序列化为[]byte，发送到服务端
func (mini MiniRpcReqCoder) Encode(ctx framework.RpcCtx, body []byte) ([]byte, error) {
	// body, _ := proto.Marshal(req)

	id := atomic.AddUint32(&reqId, 1)
	head := &miniproto.ReqHead{
		Version: 1,
		ReqId:   id,
		Api:     ctx.GetName(),
		Timeout: 0, // TODO: from ctx && config
	}

	request := &miniproto.RpcReq{
		Head: head,
		Body: body,
	}

	content, _ := proto.Marshal(request)
	return content, nil
}

// 服务端收到[]byte, 从中解析出客户端上下文及业务请求包
func (mini MiniRpcReqCoder) Decode(ctx framework.RpcCtx, data []byte) ([]byte, error) {
	var holder = &miniproto.RpcReq{}
	err := proto.Unmarshal(data, holder)

	ctx.WithName(holder.GetHead().GetApi())

	// 透传其它相关字段
	ctx.WithExt("version", holder.GetHead().GetVersion())
	ctx.WithExt("reqId", holder.GetHead().GetReqId())
	ctx.WithExt("timeout", holder.GetHead().GetTimeout())

	return holder.Body, err
}

type MiniRpcRspCoder struct {
}

// 把服务端上下文及业务响应包，序列化为[]byte，发送到客户端
func (mini MiniRpcRspCoder) Encode(ctx framework.RpcCtx, body []byte) ([]byte, error) {
	head := &miniproto.RspHead{
		Version: framework.ExtUint32Val(ctx, "version"),
		ReqId:   framework.ExtUint32Val(ctx, "reqId"),
		Ret:     framework.ExtInt32Val(ctx, "ret"),
		Msg:     []byte(framework.ExtStringVal(ctx, "msg")),
	}

	// body, _ := proto.Marshal(rsp)
	resp := &miniproto.RpcRsp{
		Head: head,
		Body: body,
	}

	return proto.Marshal(resp)
}

// 客户端收到[]byte, 从中解析出服务端上下文及业务响应包
func (mini MiniRpcRspCoder) Decode(ctx framework.RpcCtx, data []byte) ([]byte, error) {
	var holder = &miniproto.RpcRsp{}
	err := proto.Unmarshal(data, holder)

	// 透传其它相关字段
	head := holder.GetHead()
	ctx.WithExt("seq", head.GetReqId())

	if err == nil && head.GetRet() > 0 {
		err = fmt.Errorf("code:%d, message:%+v", head.Ret, string(head.Msg))
	}

	return holder.Body, err
}
