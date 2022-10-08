package core

import (
	"sync"
)

type WordManager struct {
	rw      sync.RWMutex
	AoiMgr  *AOIManager
	Players map[int32]*Player
}

var WordMgrObj *WordManager

func init() {
	WordMgrObj = &WordManager{
		AoiMgr:  NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
		Players: make(map[int32]*Player),
	}
}

func (wm *WordManager) AddPlayer(player *Player) {
	wm.rw.Lock()
	wm.Players[player.Pid] = player
	wm.rw.Unlock()

	wm.AoiMgr.AddToGridByPos(int(player.Pid), player.X, player.Z)
}

func (wm *WordManager) RemovePlayerByPid(pid int32) {
	wm.rw.Lock()
	defer wm.rw.Unlock()
	player := wm.Players[pid]
	wm.AoiMgr.RemoveFromGridByPos(int(pid), player.X, player.Z)
	delete(wm.Players, pid)
}

func (wm *WordManager) GetPlayerByPid(pid int32) *Player {
	wm.rw.RLock()
	defer wm.rw.RUnlock()
	return wm.Players[pid]
}

func (wm *WordManager) GetAllPlayer() []*Player {
	wm.rw.RLock()
	defer wm.rw.RUnlock()
	players := make([]*Player, 0)
	for _, v := range wm.Players {
		players = append(players, v)
	}
	return players
}
