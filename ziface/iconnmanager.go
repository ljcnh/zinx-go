package ziface

type IConnManager interface {
	Add(IConnection)
	Remove(IConnection)
	Get(connId uint32) (IConnection, error)
	Len() int
	ClearConn()
}
