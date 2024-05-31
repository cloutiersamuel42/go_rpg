package animation

import (
	"github.com/cloutiersamuel42/game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	_ = iota
	IdPlayerIdleAnimationUp
	IdPlayerIdleAnimationDown
	IdPlayerIdleAnimationLeft
	IdPlayerIdleAnimationRight
	IdPlayerWalkingUp
	IdPlayerWalkingDown
	IdPlayerWalkingLeft
	IdPlayerWalkingRight
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
	if a.NFrames == 1 {
		return a.Frames[0]
	}

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
	anm.RegisterAnimation(charImageAsset, []int{5}, 0, IdPlayerIdleAnimationDown)
	anm.RegisterAnimation(charImageAsset, []int{17}, 0, IdPlayerIdleAnimationLeft)
	anm.RegisterAnimation(charImageAsset, []int{29}, 0, IdPlayerIdleAnimationRight)
	anm.RegisterAnimation(charImageAsset, []int{41}, 0, IdPlayerIdleAnimationUp)
	anm.RegisterAnimation(charImageAsset, []int{4, 5, 6, 5}, 8, IdPlayerWalkingDown)
	anm.RegisterAnimation(charImageAsset, []int{16, 17, 18, 17}, 8, IdPlayerWalkingLeft)
	anm.RegisterAnimation(charImageAsset, []int{28, 29, 30, 29}, 8, IdPlayerWalkingRight)
	anm.RegisterAnimation(charImageAsset, []int{40, 41, 42, 41}, 8, IdPlayerWalkingUp)
}
