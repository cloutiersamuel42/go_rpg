package game

import (
	"math"

	"github.com/cloutiersamuel42/game/vec"
)

type Camera struct {
	Pos  vec.Vec2
	Dest vec.Vec2
}

func (c *Camera) SetDestination(v vec.Vec2) {
	c.Dest = v
}

// Move to c.Dest
func (c *Camera) MoveCamera(camSpeed float64) {
	camDX := c.Dest.X - c.Pos.X
	camDY := c.Dest.Y - c.Pos.Y

	if math.Abs(camDX) <= camSpeed && math.Abs(camDY) <= camSpeed {
		c.Pos = c.Dest
	} else {
		if camDX != 0 {
			c.Pos.X += math.Copysign(camSpeed, camDX)
		} else if camDY != 0 {
			c.Pos.Y += math.Copysign(camSpeed, camDY)
		}
	}
}
