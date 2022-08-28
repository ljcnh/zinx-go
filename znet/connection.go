package znet

import (
	"github.com/ljcnh/zinx-go/ziface"
	"log"
	"net"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnId   uint32
	isClosed bool
	ExitCh   chan bool
	Router   ziface.IRouter
}

func NewConnection(conn *net.TCPConn, ConnId uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnId:   ConnId,
		isClosed: false,
		Router:   router,
		ExitCh:   make(chan bool, 1),
	}
}

func (c *Connection) Start() {
	log.Printf("Conn Start... ConnId=%d\n", c.ConnId)
	go c.StartRead()
	//  TODO
}

func (c *Connection) StartRead() {
	log.Printf("Reading is running...\n")
	defer log.Printf("connId = %d, Reader is exit,Remote addr is: %s\n", c.ConnId, c.RemoteAddr().String())
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			log.Printf("Reading buf err: %v \n", err)
			break
		}
		req := &Request{
			conn: c,
			data: buf,
		}
		go func(req ziface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(req)
	}
}

func (c *Connection) Stop() {
	log.Printf("Conn Stop... ConnId=%d\n", c.ConnId)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	close(c.ExitCh)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
func (c *Connection) Send([]byte) error {
	return nil
}
