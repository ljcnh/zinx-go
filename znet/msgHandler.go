package znet

import (
	"github.com/ljcnh/zinx-go/ziface"
	"log"
	"strconv"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{Apis: map[uint32]ziface.IRouter{}}
}

func (m *MsgHandler) DoMsgHandler(req ziface.IRequest) {
	msgId := req.GetMsgId()
	handler, ok := m.Apis[msgId]
	if !ok {
		log.Printf("api msgId= %d, is not FOUNT! Noed Register\n", msgId)
		return
	}
	handler.PreHandle(req)
	handler.Handle(req)
	handler.PostHandle(req)
}

func (m *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgId]; ok {
		panic("repeat api, msgId= " + strconv.Itoa(int(msgId)))
	}
	m.Apis[msgId] = router
	log.Printf("Add api MsgId= %d succ\n", msgId)
}
