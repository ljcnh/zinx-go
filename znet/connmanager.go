package znet

import (
	"errors"
	"github.com/ljcnh/zinx-go/ziface"
	"log"
	"sync"
)

type ConnManager struct {
	rw          sync.RWMutex
	Connections map[uint32]ziface.IConnection
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		Connections: make(map[uint32]ziface.IConnection),
	}
}

func (cm *ConnManager) Add(conn ziface.IConnection) {
	cm.rw.Lock()
	defer cm.rw.Unlock()
	cm.Connections[conn.GetConnId()] = conn
	log.Printf("ConnId= %d add to ConnManager successfully: conn num = %d \n", conn.GetConnId(), cm.Len())
}

func (cm *ConnManager) Remove(conn ziface.IConnection) {
	cm.rw.Lock()
	defer cm.rw.Unlock()
	delete(cm.Connections, conn.GetConnId())
	log.Printf("ConnId= %d remove to ConnManager successfully: conn num = %d \n", conn.GetConnId(), cm.Len())

}

func (cm *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	cm.rw.RLock()
	defer cm.rw.RUnlock()
	if conn, ok := cm.Connections[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not FOUND")
	}
}

func (cm *ConnManager) Len() int {
	return len(cm.Connections)
}

func (cm *ConnManager) ClearConn() {
	cm.rw.Lock()
	defer cm.rw.Unlock()
	for connId, conn := range cm.Connections {
		conn.Stop()
		delete(cm.Connections, connId)
	}
	log.Printf("Clear All connections succ, conn num= %d \n", cm.Len())
}
