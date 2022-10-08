package main

import (
	"github.com/ljcnh/zinx-go/mmo-game/apis"
	"github.com/ljcnh/zinx-go/mmo-game/core"
	"github.com/ljcnh/zinx-go/ziface"
	"github.com/ljcnh/zinx-go/znet"
	"log"
)

func OnConnectionAdd(conn ziface.IConnection) {
	player := core.NewPlayer(conn)

	player.SyncPid()

	player.BroadCastStartPosition()

	core.WordMgrObj.AddPlayer(player)

	conn.SetProperty("pid", player.Pid)

	player.SyncSurrounding()

	player.LoginInfo(player.Pid)

	log.Println("Player Id ", player.Pid, " is arrived")
}

func OnConnectionLost(conn ziface.IConnection) {
	pid, _ := conn.GetProperty("pid")
	player := core.WordMgrObj.GetPlayerByPid(pid.(int32))

	player.OfflineInfo(player.Pid)

	player.Offline()
}

func main() {
	s := znet.NewServer("MMO Game")

	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	s.AddRouter(2, &apis.WordCharApi{})
	s.AddRouter(3, &apis.MoveApi{})

	s.Serve()
}
