package game

import "github.com/cloutiersamuel42/game/vec"

type Area struct {
	Name   string
	Layout []int
	ColMap []int
	TilesW int
	TilesH int
}

func (a *Area) GetCol(coords vec.Vec2) int {
	x := int(coords.X)
	y := int(coords.Y)

	// OOB
	if x >= a.TilesW || x < 0 || y >= a.TilesH || y < 0 {
		return 1
	}

	return a.ColMap[y*a.TilesH+x]
}
