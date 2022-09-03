package ziface

type IMsgHandler interface {
	DoMsgHandler(IRequest)
	AddRouter(uint32, IRouter)
	StartWorkerPool()
	SendMsgToTaskQueue(IRequest)
}
