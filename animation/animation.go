package animation

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	_ = iota
	IdPlayerIdleAnimation
)

type Animation struct {
	Frames       []*ebiten.Image
	Delay        int
	nFrames      int
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

	if a.currentFrame == a.nFrames {
		a.currentFrame = 0
	}

	fmt.Printf("On animation frame %d\n", a.currentFrame)
	return a.Frames[a.currentFrame]
}
