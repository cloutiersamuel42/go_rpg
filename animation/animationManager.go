package animation

import "fmt"

type AnimationManager struct {
	animations map[int]*Animation
}

func NewAnimationManager() *AnimationManager {
	return &AnimationManager{
		animations: make(map[int]*Animation),
	}
}

func (am *AnimationManager) RegisterAnimation(anim *Animation, id int) {
	anim.nFrames = len(anim.Frames)
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
