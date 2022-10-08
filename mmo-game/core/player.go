package core

import (
	"fmt"
	"github.com/ljcnh/zinx-go/mmo-game/pb"
	"github.com/ljcnh/zinx-go/ziface"
	"google.golang.org/protobuf/proto"
	"log"
	"math/rand"
	"sync"
)

type Player struct {
	Pid  int32
	Conn ziface.IConnection
	X    float32 // 平面的x坐标
	Y    float32 // 高度
	Z    float32 // 平面的y坐标
	V    float32 // 旋转的角度(0-360)
}

var PidGen int32 = 1
var PIDLock sync.Mutex

func NewPlayer(conn ziface.IConnection) *Player {
	PIDLock.Lock()
	defer PIDLock.Unlock()
	player := &Player{
		Pid:  PidGen,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(140 + rand.Intn(20)),
		V:    0,
	}
	PidGen++
	return player
}

func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		log.Println("Marshal msg err: ", err)
		return
	}
	if p.Conn == nil {
		log.Println("connection in player is nil")
		return
	}
	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		log.Println("Player SendMsg error!")
		return
	}
	return
}

func (p *Player) SyncPid() {
	data := &pb.SyncPid{Pid: p.Pid}
	p.SendMsg(1, data)
}

func (p *Player) BroadCastStartPosition() {
	data := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	p.SendMsg(200, data)
}

func (p *Player) Talk(content string) {
	data := &pb.BroadCast{
		Pid:  p.Pid,
		Tp:   1,
		Data: &pb.BroadCast_Content{Content: content},
	}

	players := WordMgrObj.GetAllPlayer()

	for _, p := range players {
		p.SendMsg(200, data)
	}
}

func (p *Player) SyncSurrounding() {
	pids := WordMgrObj.AoiMgr.GetPidsByPos(p.X, p.Z)
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WordMgrObj.GetPlayerByPid(int32(pid)))
	}

	data := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{P: &pb.Position{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			V: p.V,
		}},
	}

	for _, p := range players {
		p.SendMsg(200, data)
	}

	playersProto := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		p := &pb.Player{
			Pid: player.Pid,
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		playersProto = append(playersProto, p)
	}

	syncPlayersProto := &pb.SyncPlayers{Ps: playersProto[:]}
	p.SendMsg(202, syncPlayersProto)
}

func (p *Player) UpdatePos(x, y, z, v float32) {
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v

	data := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4,
		Data: &pb.BroadCast_P{P: &pb.Position{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			V: p.V,
		}},
	}

	players := p.GetSurroundingPlayers()

	for _, p := range players {
		if p == nil {
			continue
		}
		p.SendMsg(200, data)
	}
}

func (p *Player) GetSurroundingPlayers() []*Player {
	pids := WordMgrObj.AoiMgr.GetPidsByPos(p.X, p.Z)

	players := make([]*Player, 0, len(pids))

	for _, pid := range pids {
		players = append(players, WordMgrObj.GetPlayerByPid(int32(pid)))
	}
	return players
}

func (p *Player) Offline() {
	plyers := p.GetSurroundingPlayers()

	data := &pb.SyncPid{
		Pid: p.Pid,
	}

	for _, p := range plyers {
		p.SendMsg(201, data)
	}

	//WordMgrObj.AoiMgr.RemoveFromGridByPos(int(p.Pid), p.X, p.Z)
	WordMgrObj.RemovePlayerByPid(p.Pid)
}

func (p *Player) LoginInfo(pid int32) {
	loginInfo := fmt.Sprintf("player %v login", pid)

	data := &pb.Talk{Content: loginInfo}

	player := WordMgrObj.GetPlayerByPid(pid)

	player.Talk(data.Content)
}

func (p *Player) OfflineInfo(pid int32) {
	offlineInfo := fmt.Sprintf("player %v offline", pid)

	data := &pb.Talk{Content: offlineInfo}

	player := WordMgrObj.GetPlayerByPid(pid)

	player.Talk(data.Content)
}
