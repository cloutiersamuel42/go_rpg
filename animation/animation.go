package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	_ = iota
	IdPlayerIdleAnimationDown
	IdPlayerIdleAnimationUp
	IdPlayerIdleAnimationLeft
	IdPlayerIdleAnimationRight
)

type Animation struct {
	Frames       []*ebiten.Image
	Delay        int
	NFrames      int
	counter      int
	currentFrame int
}

func (a *Animation) GetCurFrame() *ebiten.Image {
	if a.counter < a.Delay {
		a.counter++
	} else {
		a.counter = 0
		a.currentFrame++
	}

	if a.currentFrame == a.NFrames {
		a.currentFrame = 0
	}

	return a.Frames[a.currentFrame]
}
