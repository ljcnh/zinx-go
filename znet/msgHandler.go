package znet

import (
	zinx_go "github.com/ljcnh/zinx-go"
	"github.com/ljcnh/zinx-go/ziface"
	"log"
	"strconv"
)

type MsgHandler struct {
	Apis           map[uint32]ziface.IRouter
	TaskQueue      []chan ziface.IRequest
	WorkerPollSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           map[uint32]ziface.IRouter{},
		WorkerPollSize: zinx_go.GlobalObject.WorkerPollSize,
		TaskQueue:      make([]chan ziface.IRequest, zinx_go.GlobalObject.WorkerPollSize),
	}
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

func (m *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPollSize); i++ {
		m.TaskQueue[i] = make(chan ziface.IRequest, zinx_go.GlobalObject.MaxWorkerTaskLen)
		go m.StartOneWorker(i, m.TaskQueue[i])
	}
}

func (m *MsgHandler) StartOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	log.Printf("WorkId=%d is started \n", workerId)
	for {
		select {
		case req := <-taskQueue:
			m.DoMsgHandler(req)
		}
	}

}

func (m *MsgHandler) SendMsgToTaskQueue(req ziface.IRequest) {
	workId := req.GetConnection().GetConnId() % m.WorkerPollSize
	log.Printf("Add ConnId= =%v, request MsgId=%v,to WorkId=%v \n", req.GetConnection().GetConnId(), req.GetMsgId(), workId)
	m.TaskQueue[workId] <- req
}
