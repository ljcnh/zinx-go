package apis

import (
	"github.com/ljcnh/zinx-go/mmo-game/core"
	"github.com/ljcnh/zinx-go/mmo-game/pb"
	"github.com/ljcnh/zinx-go/ziface"
	"github.com/ljcnh/zinx-go/znet"
	"google.golang.org/protobuf/proto"
	"log"
)

type MoveApi struct {
	znet.BaseRouter
}

func (m *MoveApi) Handle(request ziface.IRequest) {
	data := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), data)
	if err != nil {
		log.Println("Move Position Unmarshal error: ", err)
		return
	}

	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		log.Println("Move Connection GetProperty error: ", err)
		return
	}

	log.Printf("Player pid = %d, move(%f,%f,%f,%f)\n", pid, data.X, data.Y, data.Z, data.V)

	player := core.WordMgrObj.GetPlayerByPid(pid.(int32))

	player.UpdatePos(data.X, data.Y, data.Z, data.V)
}
