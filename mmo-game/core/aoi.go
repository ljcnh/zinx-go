package core

import (
	"fmt"
)

const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)

type AOIManager struct {
	MinX  int
	MaxX  int
	CntsX int
	MinY  int
	MaxY  int
	CntsY int
	grids map[int]*Grid
}

func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			gid := y*cntsX + x
			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridLength(),
				aoiMgr.MinY+(y+1)*aoiMgr.gridLength())
		}
	}
	return aoiMgr
}

func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager:\n MinX: %d, MaxX: %d, MinY: %d, MaxY: %d\n", m.MinX, m.MaxX, m.MinY, m.MaxY)
	for _, gird := range m.grids {
		s += fmt.Sprintln(gird)
	}
	return s
}

func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	if _, ok := m.grids[gID]; !ok {
		return
	}
	grids = append(grids, m.grids[gID])

	idx := gID % m.CntsX
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}
	if idx < m.CntsX-1 {
		grids = append(grids, m.grids[gID+1])
	}

	gids := make([]int, 0, len(grids))
	for _, grid := range grids {
		gids = append(gids, grid.GID)
	}
	for _, id := range gids {
		idy := id / m.CntsY
		if idy > 0 {
			grids = append(grids, m.grids[id-m.CntsX])
		}
		if idy < m.CntsY-1 {
			grids = append(grids, m.grids[id+m.CntsX])
		}
	}
	return
}

func (m *AOIManager) GetGidByPos(x, y float32) int {
	idx := (int(x) - m.MinX) / m.gridWidth()
	idy := (int(y) - m.MinY) / m.gridLength()
	return idy*m.CntsX + idx
}

func (m *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int) {
	gID := m.GetGidByPos(x, y)

	grids := m.GetSurroundGridsByGid(gID)

	for _, v := range grids {
		playerIDs = append(playerIDs, v.GetPlayerIDs()...)
	}
	return
}

func (m *AOIManager) AddPidToGrid(pID, gID int) {
	m.grids[gID].Add(pID)
}

func (m *AOIManager) RemovePidToGrid(pID, gID int) {
	m.grids[gID].Remove(pID)
}

func (m *AOIManager) GetPidsByGid(gID int) (playerIDs []int) {
	playerIDs = m.grids[gID].GetPlayerIDs()
	return
}

func (m *AOIManager) AddToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Add(pID)
}

func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Remove(pID)
}
