package main

import "github.com/ljcnh/zinx-go/znet"

func main() {
	s := znet.NewServer("zinxv0.1")
	s.Serve()
}
