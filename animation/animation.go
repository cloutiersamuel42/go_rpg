package animation

import (
	"github.com/cloutiersamuel42/game/assets"
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

// Increments animation counter and returns current image
func (a *Animation) GetCurFrame() *ebiten.Image {
	if a.counter < a.Delay {
		a.counter++
	} else {
		a.counter = 0
		a.currentFrame++
	}

	// loop animation
	if a.currentFrame == a.NFrames {
		a.currentFrame = 0
	}

	return a.Frames[a.currentFrame]
}

func InitAnimations(asm *assets.AssetManager, anm *AnimationManager) {
	charImageAsset := asm.GetAsset(assets.IdAssetCharacters).(*assets.ImageAsset)
	anm.RegisterAnimation(charImageAsset, []int{55, 56, 57, 56}, 8, IdPlayerIdleAnimationDown)
	anm.RegisterAnimation(charImageAsset, []int{67, 68, 69, 68}, 8, IdPlayerIdleAnimationLeft)
	anm.RegisterAnimation(charImageAsset, []int{79, 80, 81, 80}, 8, IdPlayerIdleAnimationRight)
	anm.RegisterAnimation(charImageAsset, []int{91, 92, 93, 92}, 8, IdPlayerIdleAnimationUp)
}
