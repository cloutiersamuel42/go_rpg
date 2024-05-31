package game

import (
	"github.com/cloutiersamuel42/game/constants"
	"github.com/cloutiersamuel42/game/vec"
	"github.com/hajimehoshi/ebiten/v2"
)

type Area struct {
	Name   string
	Layout []int
	ColMap []int
	TilesW int
	TilesH int
	NPCs   []*Character
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

func (a *Area) AddCharacter(c *Character) {
	a.NPCs = append(a.NPCs, c)
}

func (a *Area) UpdateNPCs(g *Game) {
	for _, npc := range a.NPCs {
		npc.UpdateAnimation()
		npc.UpdateCharacterLogic(g, false)
	}
}

func (a *Area) RenderNPCs(screen *ebiten.Image, cam *Camera) {
	for _, npc := range a.NPCs {
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Translate(float64(npc.Pos.X*constants.TileSize)-cam.Pos.X, float64(npc.Pos.Y*constants.TileSize)-cam.Pos.Y)
		screen.DrawImage(npc.Anim.GetCurFrame(), o)
	}
}
