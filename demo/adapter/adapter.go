package adapter

import (
	"github.com/mingo-chen/wheel-minirpc/ext"
)

func init() {
	ext.RegistCoder("minirpc", MiniRpcReqCoder{}, MiniRpcRspCoder{})
	ext.RegistFramer("minirpc", MiniRpcFramer{})
}
