package main

import (
	"fmt"
	"github.com/ljcnh/zinx-go/ziface"
	"github.com/ljcnh/zinx-go/znet"
	"log"
)

// ping test

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(req ziface.IRequest) {
	log.Printf("Call PingRouter 0...\n")
	err := req.GetConnection().SendMsg(200, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (p *HelloRouter) Handle(req ziface.IRequest) {
	log.Printf("Call HelloRouter 1...\n")
	err := req.GetConnection().SendMsg(201, []byte("hello world"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("zinxv0.5")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
