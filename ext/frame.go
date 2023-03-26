package ext

import (
	"net"
)

// Framer 帧编解码器接口定义
type Framer interface {
	ReadFrame(conn net.Conn) ([]byte, error)

	WriteFrame(conn net.Conn, data []byte) error
}

var framerHub = map[string]Framer{}

// GetFramer 获取指定协议的帧编解码器
func GetFramer(protoName string) Framer {
	return framerHub[protoName]
}

// RegistFramer 注册帧编解码器
func RegistFramer(protoName string, framer Framer) {
	framerHub[protoName] = framer
}
