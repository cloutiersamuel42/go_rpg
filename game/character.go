package game

import (
	"math"

	"github.com/cloutiersamuel42/game/animation"
	"github.com/cloutiersamuel42/game/constants"
	"github.com/cloutiersamuel42/game/vec"
	"github.com/hajimehoshi/ebiten/v2"
)

type State int
type Direction int

const (
	Idle State = iota
	MovUp
	MovDown
	MovLeft
	MovRight
)

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Character struct {
	Pos   vec.Vec2
	Dest  vec.Vec2
	Anim  *animation.Animation
	state State
	dir   Direction
	speed float64
}

func Newcharacter(initialPos vec.Vec2) *Character {
	return &Character{
		Pos:   initialPos,
		Dest:  vec.Vec2{X: 0, Y: 0},
		Anim:  nil,
		state: Idle,
		dir:   Up,
		speed: 0.10,
	}
}

// Do not check collisions here
func (c *Character) MoveCharacter() {
	dX := c.Dest.X - c.Pos.X
	dY := c.Dest.Y - c.Pos.Y

	if math.Abs(dX) <= c.speed && math.Abs(dY) <= c.speed {
		c.Pos = c.Dest
		c.state = Idle
	} else {
		if dX != 0 {
			c.Pos.X += math.Copysign(c.speed, dX)
		} else if dY != 0 {
			c.Pos.Y += math.Copysign(c.speed, dY)
		}
	}
}

func (player *Character) UpdatePlayer(g *Game, am *animation.AnimationManager) {
	if !player.Moving() {
		player.Dest = player.Pos
		g.Camera().SetDestination(g.Camera().Pos)
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			player.Dest.X -= 1
			g.Camera().Dest.X -= constants.TileSize
			player.state = MovLeft
			player.dir = Left
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.Dest.X += 1
			g.Camera().Dest.X += constants.TileSize
			player.state = MovRight
			player.dir = Right
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			player.Dest.Y -= 1
			g.Camera().Dest.Y -= constants.TileSize
			player.state = MovUp
			player.dir = Up
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			player.Dest.Y += 1
			g.Camera().Dest.Y += constants.TileSize
			player.state = MovDown
			player.dir = Down

		}
	}

	player.UpdateAnimation(am)
	player.UpdateCharacterLogic(g, true)
}

func (c *Character) UpdateAnimation(am *animation.AnimationManager) {
	// TODO put directional animations in character struct maybe
	switch c.state {
	case MovLeft:
		c.Anim = am.GetAnimation(animation.IdPlayerWalkingLeft)
	case MovRight:
		c.Anim = am.GetAnimation(animation.IdPlayerWalkingRight)
	case MovUp:
		c.Anim = am.GetAnimation(animation.IdPlayerWalkingUp)
	case MovDown:
		c.Anim = am.GetAnimation(animation.IdPlayerWalkingDown)
	case Idle:
		switch c.dir {
		case Up:
			c.Anim = am.GetAnimation(animation.IdPlayerIdleAnimationUp)
		case Down:
			c.Anim = am.GetAnimation(animation.IdPlayerIdleAnimationDown)
		case Left:
			c.Anim = am.GetAnimation(animation.IdPlayerIdleAnimationLeft)
		case Right:
			c.Anim = am.GetAnimation(animation.IdPlayerIdleAnimationRight)
		}
	}
}

func (c *Character) IsTileWalkable(g *Game, dir State) bool {
	return g.Area().GetCol(c.Dest) == 0
}

func (c *Character) UpdateCharacterLogic(g *Game, moveCam bool) {
	if c.Moving() {
		if c.IsTileWalkable(g, c.state) {
			if moveCam {
				g.Camera().MoveCamera(c.speed * constants.TileSize)
			}
			c.MoveCharacter()
		} else {
			c.state = Idle
			c.Dest = c.Pos
		}
	}
}

func (c *Character) Moving() bool {
	return c.state != Idle
}
