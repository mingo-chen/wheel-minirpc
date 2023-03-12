package core

// 编码器
// 1、把协议内容从[]byte中解析出来
// 2、把协议中接口元信息绑定到RpcCtx上，或者从RpcCtx解析出来，用于框架控制
type Coder interface {
	// 把业务数据及上下文信息打成完整的包发送给请求方
	Encode(ctx RpcCtx, body []byte) (pkg []byte, err error)

	// 把收到的包解析出业务数据及上下文信息
	Decode(ctx RpcCtx, pkg []byte) (body []byte, err error)
}

func GetCoder(protoName string) (Coder, Coder) {
	cli := coderHub[protoName+"_client"]
	svr := coderHub[protoName+"_server"]
	return cli, svr
}

var coderHub = map[string]Coder{}

func RegistCoder(protoName string, cli, svr Coder) {
	coderHub[protoName+"_client"] = cli
	coderHub[protoName+"_server"] = svr
}
