package adapter

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	framerMagic   = 0x113
	framerHeadLen = 8
)

type MiniRpcFramer struct {
}

// magic(2) + len(2) + 保留(4) + pkg(len)
func (mini MiniRpcFramer) ReadFrame(conn net.Conn) ([]byte, error) {
	var header [framerHeadLen]byte
	n, err := io.ReadFull(conn, header[:])
	if err != nil {
		return nil, err
	}
	if n != framerHeadLen {
		return nil, fmt.Errorf("invalid framer header, want:%d, actual:%d", framerHeadLen, n)
	}

	magic := binary.BigEndian.Uint16(header[0:2])
	if magic != framerMagic {
		return nil, fmt.Errorf("invalid framer, want:%d, actual:%d", framerMagic, magic)
	}

	totalLen := binary.BigEndian.Uint16(header[2:4])
	var pkgLen = int(totalLen) - framerHeadLen
	var pkg = make([]byte, pkgLen)
	n, err = io.ReadFull(conn, pkg)
	if err != nil {
		return nil, err
	}
	if n != pkgLen {
		return nil, fmt.Errorf("invalid framer pkg, want:%d, actual:%d", pkgLen, n)
	}

	// copy(pkg, header[:])
	return pkg, nil
}

func (mini MiniRpcFramer) WriteFrame(conn net.Conn, data []byte) error {
	var header [framerHeadLen]byte
	binary.BigEndian.PutUint16(header[:2], framerMagic)
	totalLen := framerHeadLen + len(data)
	binary.BigEndian.PutUint16(header[2:4], uint16(totalLen))
	binary.BigEndian.PutUint32(header[4:framerHeadLen], 0)

	n, err := conn.Write(append(header[:], data...))
	if err != nil {
		return err
	}

	if n != totalLen {
		return fmt.Errorf("invalid framer, want:%d, actual:%d", totalLen, n)
	}

	return nil
}
