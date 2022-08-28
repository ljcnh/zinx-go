package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("client start...")
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8900")
	if err != nil {
		log.Printf("client start err: %v\n", err)
		return
	}
	for {
		_, err := conn.Write([]byte("Hello zinx"))
		if err != nil {
			log.Printf("write conn err: %v\n", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Printf("read buf err: %v\n", err)
			return
		}
		log.Printf("server call back: %s,cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}
