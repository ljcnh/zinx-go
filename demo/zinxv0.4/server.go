package main

import (
	"github.com/ljcnh/zinx-go/ziface"
	"github.com/ljcnh/zinx-go/znet"
	"log"
)

// ping test

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) PreHandle(req ziface.IRequest) {
	log.Printf("Call Router PreHandle...\n")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("Before ping...\n"))
	if err != nil {
		log.Printf("Call Back BeforeHandle error...\n")
	}
}

func (p *PingRouter) Handle(req ziface.IRequest) {
	log.Printf("Call Router Handle...\n")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping...\n"))
	if err != nil {
		log.Printf("Call Back Handle error...\n")
	}
}

func (p *PingRouter) PostHandle(req ziface.IRequest) {
	log.Printf("Call Router PostHandle...\n")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("After ping...\n"))
	if err != nil {
		log.Printf("Call Back PostHandle error...\n")
	}
}

func main() {
	s := znet.NewServer("zinxv0.3")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
