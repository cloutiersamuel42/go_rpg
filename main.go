package main

import (
	_ "image/png"
	"log"

	"github.com/cloutiersamuel42/game/animation"
	"github.com/cloutiersamuel42/game/assets"
	"github.com/cloutiersamuel42/game/constants"
	"github.com/cloutiersamuel42/game/player"
	"github.com/cloutiersamuel42/game/vec"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	gAssetManager     *assets.AssetManager        = assets.NewAssetManager()
	gAnimationManager *animation.AnimationManager = animation.NewAnimationManager()
)

type Game struct {
	area   Area
	cam    Camera
	player player.Player
}

type Camera struct {
	pos vec.Vec2
}

type Area struct {
	name   string
	layout []int
	tilesW int
	tilesH int
}

func init() {
	assets.InitAssets(gAssetManager)

	charImageAsset := gAssetManager.GetAsset(assets.IdAssetCharacters).(*assets.ImageAsset)
	gAnimationManager.RegisterAnimation(charImageAsset, []int{55, 56, 57, 56}, 8, animation.IdPlayerIdleAnimationDown)
	gAnimationManager.RegisterAnimation(charImageAsset, []int{67, 68, 69, 68}, 8, animation.IdPlayerIdleAnimationLeft)
	gAnimationManager.RegisterAnimation(charImageAsset, []int{79, 80, 81, 80}, 8, animation.IdPlayerIdleAnimationRight)
	gAnimationManager.RegisterAnimation(charImageAsset, []int{91, 92, 93, 92}, 8, animation.IdPlayerIdleAnimationUp)
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.cam.pos.X += constants.TileSize
		g.player.Pos.X -= 1
		g.player.Anim = gAnimationManager.GetAnimation(animation.IdPlayerIdleAnimationLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.cam.pos.X -= constants.TileSize
		g.player.Pos.X += 1
		g.player.Anim = gAnimationManager.GetAnimation(animation.IdPlayerIdleAnimationRight)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.cam.pos.Y += constants.TileSize
		g.player.Pos.Y -= 1
		g.player.Anim = gAnimationManager.GetAnimation(animation.IdPlayerIdleAnimationUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.cam.pos.Y -= constants.TileSize
		g.player.Pos.Y += 1
		g.player.Anim = gAnimationManager.GetAnimation(animation.IdPlayerIdleAnimationDown)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tilesAsset := gAssetManager.GetAsset(assets.IdAssetBasictiles).(*assets.ImageAsset)
	for i := range g.area.layout {
		x := i % g.area.tilesW
		y := i / g.area.tilesW
		if g.area.layout[i] != 0 {
			o := &ebiten.DrawImageOptions{}
			o.GeoM.Translate(g.cam.pos.X+float64(x*constants.TileSize), g.cam.pos.Y+float64(y*constants.TileSize))
			screen.DrawImage(tilesAsset.GetTileFromOffset(g.area.layout[i]), o)
		}
		// fmt.Printf("Drawing tile {x:%d y:%d}\n", x, y)
	}

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Translate(g.cam.pos.X+float64(g.player.Pos.X*constants.TileSize), g.cam.pos.Y+float64(g.player.Pos.Y*constants.TileSize))
	screen.DrawImage(g.player.Anim.GetCurFrame(), o)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	g := &Game{
		player: player.Player{
			Pos:  vec.Vec2{X: 4, Y: 4},
			Anim: gAnimationManager.GetAnimation(animation.IdPlayerIdleAnimationDown),
		},
		cam: Camera{
			pos: vec.Vec2{X: 0, Y: 0},
		},
		area: Area{
			name: "Test area",
			layout: []int{
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
				66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
			},
			tilesW: 10,
			tilesH: 10,
		},
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
