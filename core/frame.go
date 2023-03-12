package core

import (
	"net"
)

// Framer 各种协议帧的读写操作
type Framer interface {
	ReadFrame(conn net.Conn) ([]byte, error)

	WriteFrame(conn net.Conn, data []byte) error
}

var framerHub = map[string]Framer{}

func GetFramer(protoName string) Framer {
	return framerHub[protoName]
}

func RegistFramer(protoName string, framer Framer) {
	framerHub[protoName] = framer
}
