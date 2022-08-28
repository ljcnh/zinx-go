package zinx_go

import (
	"encoding/json"
	"github.com/ljcnh/zinx-go/ziface"
	"io/ioutil"
)

type GlobalObj struct {
	// server
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string
	// zinx
	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	// default
	GlobalObject = &GlobalObj{
		Host:           "0.0.0.0",
		TcpPort:        8999,
		Name:           "ZinxServer",
		Version:        "v0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	GlobalObject.Reload()
}
