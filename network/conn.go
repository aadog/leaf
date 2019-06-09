package network

import (
	"net"
)

type Conn interface {
	Conn()
	DeConn()
	ReadMsg() ([]byte, error)
	WriteMsg(args ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}
