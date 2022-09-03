package znet

import (
	"fmt"
	"github.com/ljcnh/zinx-go"
	"github.com/ljcnh/zinx-go/ziface"
	"log"
	"net"
)

type Server struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int
	MsgHandler  ziface.IMsgHandler
	ConnMgr     ziface.IConnManager
	OnConnStart func(conn ziface.IConnection)
	OnConnStop  func(conn ziface.IConnection)
}

func (s *Server) Start() {
	log.Printf("[Config] Server Name: %s, listener at IP: %s, Port: %d is starting\n", zinx_go.GlobalObject.Name, zinx_go.GlobalObject.Host, zinx_go.GlobalObject.TcpPort)
	log.Printf("Version: %s, MaxConn: %d, MaxPackageSize: %d\n", zinx_go.GlobalObject.Version, zinx_go.GlobalObject.MaxConn, zinx_go.GlobalObject.MaxPackageSize)
	go func() {
		s.MsgHandler.StartWorkerPool()

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Printf("resolve tcp addr error: %v\n", err)
			return
		}
		lis, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Printf("listen: %v, err: %v\n", s.IPVersion, err)
			return
		}
		log.Printf("start Zinx server success: %v, success, Listenning...\n", s.Name)

		var id uint32 = 0
		for {
			conn, err := lis.AcceptTCP()
			if err != nil {
				log.Printf("Accept err: %v\n", err)
				continue
			}

			if s.ConnMgr.Len() >= zinx_go.GlobalObject.MaxConn {
				log.Printf("Too Many Connection! The MaxConn= %d\n", zinx_go.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			connection := NewConnection(s, conn, id, s.MsgHandler)
			id++
			go connection.Start()
		}
	}()
}

func (s *Server) Stop() {
	log.Printf("[Stop] Server]\n")
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()
	// TODO

	// 阻塞
	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	log.Printf("AddRouter success...\n")
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       zinx_go.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         zinx_go.GlobalObject.Host,
		Port:       zinx_go.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

func (s *Server) SetOnConnStart(f func(connection ziface.IConnection)) {
	s.OnConnStart = f
}

func (s *Server) SetOnConnStop(f func(connection ziface.IConnection)) {
	s.OnConnStop = f
}

func (s *Server) CallConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		log.Printf("————> Call OnconnStart() \n")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		log.Printf("————> Call CallConnStop() \n")
		s.OnConnStop(conn)
	}
}
