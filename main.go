package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/cloutiersamuel42/game/animation"
	"github.com/cloutiersamuel42/game/assets"
	"github.com/cloutiersamuel42/game/camera"
	"github.com/cloutiersamuel42/game/character"
	"github.com/cloutiersamuel42/game/constants"
	"github.com/cloutiersamuel42/game/vec"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	gAssetManager     *assets.AssetManager        = assets.NewAssetManager()
	gAnimationManager *animation.AnimationManager = animation.NewAnimationManager()
)

type Game struct {
	area   Area
	cam    camera.Camera
	player character.Character
}

func (g *Game) Camera() *camera.Camera {
	return &g.cam
}

func (g *Game) Player() *character.Character {
	return &g.player
}

type Area struct {
	name   string
	layout []int
	tilesW int
	tilesH int
}

func init() {
	assets.InitAssets(gAssetManager)
	animation.InitAnimations(gAssetManager, gAnimationManager)
}

func (g *Game) Update() error {
	g.Player().UpdatePlayer(g, gAnimationManager)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tilesAsset := gAssetManager.GetAsset(assets.IdAssetBasictiles).(*assets.ImageAsset)
	for i := range g.area.layout {
		x := i % g.area.tilesW
		y := i / g.area.tilesW
		if g.area.layout[i] != 0 {
			o := &ebiten.DrawImageOptions{}
			o.GeoM.Translate(g.Camera().Pos.X+float64(x*constants.TileSize), g.Camera().Pos.Y+float64(y*constants.TileSize))
			screen.DrawImage(tilesAsset.GetTileFromOffset(g.area.layout[i]), o)
		}
		// fmt.Printf("Drawing tile {x:%d y:%d}\n", x, y)
	}

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Translate(g.Camera().Pos.X+float64(g.Player().Pos.X*constants.TileSize), g.Camera().Pos.Y+float64(g.Player().Pos.Y*constants.TileSize))
	screen.DrawImage(g.Player().Anim.GetCurFrame(), o)

	// Debug print
	strPos := fmt.Sprintf("Player pos => x: %f y: %f", g.Player().Pos.X, g.Player().Pos.Y)
	strDest := fmt.Sprintf("Player dir => x: %f y: %f", g.Player().Dest.X, g.Player().Dest.Y)
	ebitenutil.DebugPrintAt(screen, strPos, 0, 0)
	ebitenutil.DebugPrintAt(screen, strDest, 0, 20)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	g := &Game{
		player: *character.Newcharacter(vec.Vec2{X: 2, Y: 2}),
		cam: camera.Camera{
			Pos: vec.Vec2{X: 0, Y: 0},
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
	g.Player().Anim = gAnimationManager.GetAnimation(animation.IdPlayerIdleAnimationDown)
	ebiten.SetWindowSize(1080, 960)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
