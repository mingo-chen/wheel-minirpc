package adapter

import "github.com/mingo-chen/wheel-minirpc/core"

func init() {
	core.RegistCoder("minirpc", MiniRpcReqCoder{}, MiniRpcRspCoder{})
	core.RegistFramer("minirpc", MiniRpcFramer{})
}
