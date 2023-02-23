package ext

import (
	"testing"

	"github.com/mingo-chen/wheel-minirpc/framework"
	miniproto "github.com/mingo-chen/wheel-minirpc/framework/ext/proto"
	"google.golang.org/protobuf/proto"
)

func TestMiniRpcReqCoder_Encode(t *testing.T) {
	var coder framework.Coder = MiniRpcReqCoder{}

	req := miniproto.RpcReq{
		Head: &miniproto.ReqHead{
			Version: 1,
			ReqId:   10086,
			Api:     "get_user_by_id",
			Timeout: 2000,
		},
		Body: []byte("Hello, mini rpc!"),
	}

	ctx := framework.NewDefaultRpcCtx()

	body, _ := proto.Marshal(&req)
	data, err := coder.Encode(ctx, body)
	if err != nil {
		t.Logf("encode err:%+v", err)
	}

	after, err := coder.Decode(ctx, data)
	if err != nil {
		t.Logf("decode err:%+v", err)
	}

	var rsp miniproto.RpcRsp
	proto.Unmarshal(after, &rsp)
	t.Logf("after codec: %+v", &rsp)
}
