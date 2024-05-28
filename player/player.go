package player

import (
	"github.com/cloutiersamuel42/game/animation"
	"github.com/cloutiersamuel42/game/vec"
)

type Player struct {
	Pos  vec.Vec2
	Anim *animation.Animation
}
