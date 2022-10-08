package apis

import (
	"github.com/ljcnh/zinx-go/mmo-game/core"
	"github.com/ljcnh/zinx-go/mmo-game/pb"
	"github.com/ljcnh/zinx-go/ziface"
	"github.com/ljcnh/zinx-go/znet"
	"google.golang.org/protobuf/proto"
	"log"
)

type WordCharApi struct {
	znet.BaseRouter
}

func (wc *WordCharApi) Handle(request ziface.IRequest) {
	data := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), data)
	if err != nil {
		log.Println("Talk Unmarshal error: ", err)
		return
	}

	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		log.Println("Talk Connection GetProperty error: ", err)
		return
	}
	player := core.WordMgrObj.GetPlayerByPid(pid.(int32))

	player.Talk(data.Content)
}
