package main

import (
	"fmt"
	"github.com/ljcnh/zinx-go/znet"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("client0 start...")
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8900")
	if err != nil {
		log.Printf("client start err: %v\n", err)
		return
	}
	for {
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMessage(1, []byte("ZinxV0.6 client Test")))
		if err != nil {
			log.Printf("pack error: %v\n", err)
			return
		}
		_, err = conn.Write(binaryMsg)
		if err != nil {
			log.Printf("conn write error: %v\n", err)
			return
		}

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
			log.Printf("read head error: %v\n", err)
			return
		}
		msg, err := dp.UnPack(headData)
		if err != nil {
			log.Printf("UnPack msg error: %v\n", err)
			return
		}

		if msg.GetMsgLen() > 0 {
			data := make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, data); err != nil {
				log.Printf("read msg data error: %v\n", err)
				return
			}
			msg.SetData(data)

			fmt.Println("————> Recv Server Msg Id:", msg.GetMsgId(), " len= ", msg.GetMsgLen(), " data= ", string(msg.GetData()))
		}

		time.Sleep(1 * time.Second)
	}
}
