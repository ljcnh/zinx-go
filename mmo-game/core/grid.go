package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	GID       int
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	playerIDs map[int]bool
	rw        sync.RWMutex
}

func NewGrid(GID, MinX, MaxX, MinY, MaxY int) *Grid {
	return &Grid{
		GID:       GID,
		MinX:      MinX,
		MaxX:      MaxX,
		MinY:      MinY,
		MaxY:      MaxY,
		playerIDs: make(map[int]bool),
	}
}

func (g *Grid) Add(playerId int) {
	g.rw.Lock()
	defer g.rw.Unlock()
	g.playerIDs[playerId] = true
}

func (g *Grid) Remove(playerId int) {
	g.rw.Lock()
	defer g.rw.Unlock()
	delete(g.playerIDs, playerId)
}

func (g *Grid) GetPlayerIDs() (playerIDs []int) {
	g.rw.RLock()
	defer g.rw.RUnlock()
	for k, _ := range g.playerIDs {
		playerIDs = append(playerIDs, k)
	}
	return
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid id: %d, minX: %d, maxX: %d,minY: %d,maxY:: %d,playerIDs: %v,", g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs)
}
