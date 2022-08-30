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
	log.Printf("Call Router Handle...\n")
	err := req.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("zinxv0.5")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
