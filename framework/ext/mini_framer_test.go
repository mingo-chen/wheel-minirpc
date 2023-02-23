package ext

import (
	"bytes"
	"net"
	"testing"
	"time"
)

type myconn struct {
	stream *bytes.Buffer
}

func (conn myconn) Read(b []byte) (n int, err error) {
	return conn.stream.Read(b)
}

func (conn myconn) Write(b []byte) (n int, err error) {
	return conn.stream.Write(b)
}

func (conn myconn) Close() error {
	return nil
}

func (conn myconn) LocalAddr() net.Addr {
	return nil
}

func (conn myconn) RemoteAddr() net.Addr {
	return nil
}

func (conn myconn) SetDeadline(t time.Time) error {
	return nil
}

func (conn myconn) SetReadDeadline(t time.Time) error {
	return nil
}

func (conn myconn) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestMiniRpcFramer(t *testing.T) {
	conn := myconn{stream: bytes.NewBufferString("")}
	framer := MiniRpcFramer{}

	err := framer.WriteFrame(conn, []byte("Hello world! rtyuh"))
	if err != nil {
		t.Fatalf("write err:%+v", err)
	}

	content, err := framer.ReadFrame(conn)
	if err != nil {
		t.Fatalf("read err:%+v", err)
	}

	t.Logf("--> %+v", string(content))
}
