package adapter

import (
	"fmt"
	"sync/atomic"

	"github.com/mingo-chen/wheel-minirpc/core"
	miniproto "github.com/mingo-chen/wheel-minirpc/demo/adapter/proto"
	"google.golang.org/protobuf/proto"
)

type MiniRpcReqCoder struct {
}

var reqId uint32

// 把请求上下文及业务请求包，序列化为[]byte，发送到服务端
func (mini MiniRpcReqCoder) Encode(ctx core.RpcCtx, body []byte) ([]byte, error) {
	// body, _ := proto.Marshal(req)

	id := atomic.AddUint32(&reqId, 1)
	head := &miniproto.ReqHead{
		Version: 1,
		Seq:     id,
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
func (mini MiniRpcReqCoder) Decode(ctx core.RpcCtx, data []byte) ([]byte, error) {
	var holder = &miniproto.RpcReq{}
	err := proto.Unmarshal(data, holder)

	ctx.WithName(holder.GetHead().GetApi())

	// 透传其它相关字段
	ctx.WithExt("version", holder.GetHead().GetVersion())
	ctx.WithExt("seq", holder.GetHead().GetSeq())
	ctx.WithExt("timeout", holder.GetHead().GetTimeout())

	return holder.Body, err
}

type MiniRpcRspCoder struct {
}

// 把服务端上下文及业务响应包，序列化为[]byte，发送到客户端
func (mini MiniRpcRspCoder) Encode(ctx core.RpcCtx, body []byte) ([]byte, error) {
	head := &miniproto.RspHead{
		Version: core.ExtUint32Val(ctx, "version"),
		Seq:     core.ExtUint32Val(ctx, "seq"),
		Ret:     core.ExtInt32Val(ctx, "ret"),
		Msg:     []byte(core.ExtStringVal(ctx, "msg")),
	}

	// body, _ := proto.Marshal(rsp)
	resp := &miniproto.RpcRsp{
		Head: head,
		Body: body,
	}

	return proto.Marshal(resp)
}

// 客户端收到[]byte, 从中解析出服务端上下文及业务响应包
func (mini MiniRpcRspCoder) Decode(ctx core.RpcCtx, data []byte) ([]byte, error) {
	var holder = &miniproto.RpcRsp{}
	err := proto.Unmarshal(data, holder)

	// 透传其它相关字段
	head := holder.GetHead()
	ctx.WithExt("seq", head.GetSeq())

	if err == nil && head.GetRet() > 0 {
		err = fmt.Errorf("code:%d, message:%+v", head.Ret, string(head.Msg))
	}

	return holder.Body, err
}
