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

func ConnStart(conn ziface.IConnection) {
	log.Printf("————> ConnStart is Called...\n")
	if err := conn.SendMsg(202, []byte("ConnStart Start")); err != nil {
		log.Println(err)
	}

	conn.SetProperty("Name", "Lj")
	conn.SetProperty("Github", "github.ljcnh.com")
}

func ConnStop(conn ziface.IConnection) {
	log.Printf("————> ConnStop is Called...\n")
	log.Println("connId = ", conn.GetConnId(), "is Lost...")

	if name, err := conn.GetProperty("Name"); err == nil {
		log.Println("Name = ", name)
	}
	if github, err := conn.GetProperty("Github"); err == nil {
		log.Println("Github = ", github)
	}
}

func main() {
	s := znet.NewServer("zinxv0.9")
	s.SetOnConnStart(ConnStart)
	s.SetOnConnStop(ConnStop)
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
