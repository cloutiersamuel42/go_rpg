package animation

import (
	"fmt"

	"github.com/cloutiersamuel42/game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationManager struct {
	animations map[int]*Animation
}

func NewAnimationManager() *AnimationManager {
	return &AnimationManager{
		animations: make(map[int]*Animation),
	}
}

// Register an animation to an ID using list of tiles in image and a delay (frames per image)
func (am *AnimationManager) RegisterAnimation(img *assets.ImageAsset, frameOffsets []int, delay int, id int) {
	len := len(frameOffsets)
	frames := make([]*ebiten.Image, len)

	for i, frameOffset := range frameOffsets {
		frames[i] = img.GetTileFromOffset(frameOffset)
	}

	anim := &Animation{
		Frames:  frames,
		Delay:   delay,
		NFrames: len,
	}
	am.animations[id] = anim
	fmt.Printf("Registered animation: %d\n", id)
}

func (am *AnimationManager) GetAnimation(id int) *Animation {
	anim, ok := am.animations[id]
	if !ok {
		return nil
	}
	return anim
}
