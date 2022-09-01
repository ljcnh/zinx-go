package znet

import (
	"errors"
	"github.com/ljcnh/zinx-go/ziface"
	"io"
	"log"
	"net"
)

type Connection struct {
	Conn       *net.TCPConn
	ConnId     uint32
	isClosed   bool
	ExitCh     chan bool
	msgCh      chan []byte
	MsgHandler ziface.IMsgHandler
}

func NewConnection(conn *net.TCPConn, ConnId uint32, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:       conn,
		ConnId:     ConnId,
		isClosed:   false,
		MsgHandler: msgHandler,
		msgCh:      make(chan []byte),
		ExitCh:     make(chan bool, 1),
	}
}

func (c *Connection) Start() {
	log.Printf("Conn Start... ConnId=%d\n", c.ConnId)
	go c.StartRead()
	go c.StartWriter()
}

func (c *Connection) StartRead() {
	log.Printf("Reading Gortine is running...\n")
	defer log.Printf("Reader is exit, connId = %d, Remote addr is: %s\n", c.ConnId, c.RemoteAddr().String())
	defer c.Stop()
	for {

		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			log.Printf("Read msg head data err: %v\n", err)
			break
		}

		msg, err := dp.UnPack(headData)
		if err != nil {
			log.Printf("UnPack msg data err: %v\n", err)
			break
		}
		data := make([]byte, 0)
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				log.Printf("Recv msg data err: %v\n", err)
				break
			}
		}
		msg.SetData(data)

		req := &Request{
			conn: c,
			msg:  msg,
		}

		go c.MsgHandler.DoMsgHandler(req)
	}
}

func (c *Connection) StartWriter() {
	log.Printf("Write Gortine is running...\n")
	defer log.Printf("[conn Writer exit]: %s \n", c.RemoteAddr().String())
	for {
		select {
		case data := <-c.msgCh:
			if _, err := c.Conn.Write(data); err != nil {
				log.Printf("Send data error: %v \n", err)
				return
			}
		case <-c.ExitCh:
			return
		}
	}
}

func (c *Connection) Stop() {
	log.Printf("Conn Stop... ConnId=%d\n", c.ConnId)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	c.ExitCh <- true
	close(c.msgCh)
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection closed when sen msg")
	}

	dp := NewDataPack()

	msg := NewMessage(msgId, data)

	finalData, err := dp.Pack(msg)
	if err != nil {
		log.Printf("pack error msg id= %d\n", msgId)
		return errors.New("pack error msg")
	}
	c.msgCh <- finalData
	return nil
}
