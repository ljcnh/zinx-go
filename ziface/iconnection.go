package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnId() uint32
	RemoteAddr() net.Addr
	Send([]byte) error
}

// HandleFunc TcpConn 内容 处理的长度
type HandleFunc func(*net.TCPConn, []byte, int) error
