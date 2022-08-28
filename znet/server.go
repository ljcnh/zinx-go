package znet

import (
	"fmt"
	"github.com/ljcnh/zinx-go/ziface"
	"log"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    ziface.IRouter
}

func (s *Server) Start() {
	log.Printf("[Start] Server Listenner at IP: %v, Port %d is starting\n", s.IP, &s.Port)
	go func() {
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

			connection := NewConnection(conn, id, s.Router)
			id++
			go connection.Start()
		}
	}()
}

func (s *Server) Stop() {
	// TODO
}

func (s *Server) Serve() {
	s.Start()
	// TODO

	// 阻塞
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	log.Printf("AddRouter success...\n")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
