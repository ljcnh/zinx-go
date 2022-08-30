package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("server accept error: ", err)
				return
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head data error: ", err)
						break
					}
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("unpack head data error: ", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data err: ", err)
							return
						}
						fmt.Println("————>Recv MagId: ", msg.Id, " DataLen= ", msg.DataLen, " Data=", string(msg.Data))
					}
				}
			}(conn)

		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err: ", err)
		return
	}
	dp := NewDataPack()
	sendMsg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', '0'},
	}
	senData1, err := dp.Pack(sendMsg1)
	if err != nil {
		fmt.Println("data1 pack err: ", err)
		return
	}

	sendMsg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'W', 'o', 'r', 'd', '.', '!', '!'},
	}
	senData2, err := dp.Pack(sendMsg2)
	if err != nil {
		fmt.Println("data2 pack err: ", err)
		return
	}
	senData1 = append(senData1, senData2...)

	conn.Write(senData1)

	conn.Close()

	select {}
}
