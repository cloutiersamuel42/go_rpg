package main

import (
	"log"

	"github.com/cloutiersamuel42/game/animation"
	"github.com/cloutiersamuel42/game/game"
	"github.com/cloutiersamuel42/game/vec"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	g.Area().AddCharacter(game.Newcharacter(vec.Vec2{X: 0, Y: 0}))
	g.Area().NPCs[0].Animations = game.CharacterDirectionAnimation{
		IdleUp:    game.AnimationManager.GetAnimation(animation.IdPlayerIdleAnimationUp),
		IdleDown:  game.AnimationManager.GetAnimation(animation.IdPlayerIdleAnimationDown),
		IdleLeft:  game.AnimationManager.GetAnimation(animation.IdPlayerIdleAnimationLeft),
		IdleRight: game.AnimationManager.GetAnimation(animation.IdPlayerIdleAnimationRight),
		WalkUp:    game.AnimationManager.GetAnimation(animation.IdPlayerWalkingUp),
		WalkDown:  game.AnimationManager.GetAnimation(animation.IdPlayerWalkingDown),
		WalkLeft:  game.AnimationManager.GetAnimation(animation.IdPlayerWalkingLeft),
		WalkRight: game.AnimationManager.GetAnimation(animation.IdPlayerWalkingRight),
	}
	g.Area().NPCs[0].UpdateAnimation()
	// Make prefabs for this later
	g.Player().Animations = game.CharacterDirectionAnimation{
		IdleUp:    game.AnimationManager.GetAnimation(animation.IdPlayerIdleAnimationUp),
		IdleDown:  game.AnimationManager.GetAnimation(animation.IdPlayerIdleAnimationDown),
		IdleLeft:  game.AnimationManager.GetAnimation(animation.IdPlayerIdleAnimationLeft),
		IdleRight: game.AnimationManager.GetAnimation(animation.IdPlayerIdleAnimationRight),
		WalkUp:    game.AnimationManager.GetAnimation(animation.IdPlayerWalkingUp),
		WalkDown:  game.AnimationManager.GetAnimation(animation.IdPlayerWalkingDown),
		WalkLeft:  game.AnimationManager.GetAnimation(animation.IdPlayerWalkingLeft),
		WalkRight: game.AnimationManager.GetAnimation(animation.IdPlayerWalkingRight),
	}
	g.Player().UpdateAnimation()
	ebiten.SetWindowSize(320*2, 240*2)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
