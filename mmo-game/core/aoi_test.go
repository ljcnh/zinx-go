package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	aoimgr := NewAOIManager(100, 300, 4, 200, 450, 5)
	fmt.Println(aoimgr)
}

func TestAOIManager_GetSurroundGridsByGid(t *testing.T) {
	aoimgr := NewAOIManager(0, 250, 5, 0, 250, 5)

	for gid, _ := range aoimgr.grids {
		fmt.Println("Gid", gid)
		grids := aoimgr.GetSurroundGridsByGid(gid)
		gIDs := make([]int, 0, len(grids))
		for _, grid := range grids {
			gIDs = append(gIDs, grid.GID)
		}
		fmt.Println(gIDs)
	}
}
