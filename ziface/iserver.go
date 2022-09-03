package ziface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(uint32, IRouter)
	GetConnMgr() IConnManager
	SetOnConnStart(func(IConnection))
	SetOnConnStop(func(IConnection))
	CallConnStart(IConnection)
	CallConnStop(IConnection)
}
