package character

import (
	"math"

	"github.com/cloutiersamuel42/game/animation"
	"github.com/cloutiersamuel42/game/camera"
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

type GameInterface interface {
	Camera() *camera.Camera
	Player() *Character
}

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

func (player *Character) UpdatePlayer(g GameInterface, am *animation.AnimationManager) {
	if !player.Moving() {
		player.Dest = player.Pos
		g.Camera().SetDestination(g.Camera().Pos)
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			player.Dest.X -= 1
			g.Camera().Dest.X -= constants.TileSize
			player.state = MovLeft
			player.dir = Left
			player.UpdateAnimation(am)
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.Dest.X += 1
			g.Camera().Dest.X += constants.TileSize
			player.state = MovRight
			player.dir = Right
			player.UpdateAnimation(am)
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			player.Dest.Y -= 1
			g.Camera().Dest.Y -= constants.TileSize
			player.state = MovUp
			player.dir = Up
			player.UpdateAnimation(am)
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			player.Dest.Y += 1
			g.Camera().Dest.Y += constants.TileSize
			player.state = MovDown
			player.dir = Down
			player.UpdateAnimation(am)
		}
	}

	if player.Moving() {
		dX := player.Dest.X - player.Pos.X
		dY := player.Dest.Y - player.Pos.Y

		g.Camera().MoveCamera(player.speed * constants.TileSize)

		if math.Abs(dX) <= player.speed && math.Abs(dY) <= player.speed {
			player.Pos = player.Dest
			player.state = Idle
		} else {
			if dX != 0 {
				player.Pos.X += math.Copysign(player.speed, dX)
			} else if dY != 0 {
				player.Pos.Y += math.Copysign(player.speed, dY)
			}
		}
	}
}

func (c *Character) UpdateAnimation(am *animation.AnimationManager) {
	// TODO put directional animations in character struct maybe
	switch c.state {
	case MovLeft:
		c.Anim = am.GetAnimation(animation.IdPlayerIdleAnimationLeft)
	case MovRight:
		c.Anim = am.GetAnimation(animation.IdPlayerIdleAnimationRight)
	case MovUp:
		c.Anim = am.GetAnimation(animation.IdPlayerIdleAnimationUp)
	case MovDown:
		c.Anim = am.GetAnimation(animation.IdPlayerIdleAnimationDown)
	}
}

func (c *Character) UpdateCharacterLogic() {
	switch c.state {
	case MovLeft:

	}
}

func (c *Character) Moving() bool {
	return c.state != Idle
}
