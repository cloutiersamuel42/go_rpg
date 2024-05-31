package main

import (
	"log"

	"github.com/cloutiersamuel42/game/animation"
	"github.com/cloutiersamuel42/game/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	g.Player().Anim = game.AnimationManager.GetAnimation(animation.IdPlayerIdleAnimationDown)
	ebiten.SetWindowSize(320*2, 240*2)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
