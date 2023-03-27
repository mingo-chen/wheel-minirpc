package ext

import "github.com/mingo-chen/wheel-minirpc/core"

// 编码器
// 1、把协议内容从[]byte中解析出来
// 2、把协议中接口元信息绑定到RpcCtx上，或者从RpcCtx解析出来，用于框架控制
type Coder interface {
	// 把业务数据及上下文信息打成完整的包发送给请求方
	Encode(ctx core.RpcCtx, body []byte) (pkg []byte, err error)

	// 把收到的包解析出业务数据及上下文信息
	Decode(ctx core.RpcCtx, pkg []byte) (body []byte, err error)
}

var cliCoderHub = map[string]Coder{}
var svrCoderHub = map[string]Coder{}

// GetCoder 获取编解码器
func GetCoder(protoName string) (Coder, Coder) {
	cli := cliCoderHub[protoName]
	svr := svrCoderHub[protoName]
	return cli, svr
}

// RegistCoder 注册编解码器
func RegistCoder(protoName string, cli, svr Coder) {
	cliCoderHub[protoName] = cli
	svrCoderHub[protoName] = svr
}
