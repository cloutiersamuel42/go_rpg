package main

import (
	_ "image/png"
	"log"

	"github.com/cloutiersamuel42/game/animation"
	"github.com/cloutiersamuel42/game/assets"
	"github.com/cloutiersamuel42/game/constants"
	"github.com/cloutiersamuel42/game/vec"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var gAssetManager *assets.AssetManager
var gAnimationManager *animation.AnimationManager

type Game struct {
	area   Area
	cam    Camera
	player Player
}

type Camera struct {
	pos vec.Vec2
}

type Player struct {
	pos  vec.Vec2
	anim *animation.Animation
}

type Area struct {
	name   string
	layout []int
	tilesW int
	tilesH int
}

func init() {
	gAssetManager = assets.NewAssetManager()
	gAssetManager.RegisterAsset(&assets.ImageAsset{Path: "data/basictiles.png"}, assets.IdAssetBasictiles)
	gAssetManager.RegisterAsset(&assets.ImageAsset{Path: "data/characters.png"}, assets.IdAssetCharacters)
	gAssetManager.LoadAssets()

	charImageAsset := gAssetManager.GetAsset(assets.IdAssetCharacters).(*assets.ImageAsset)
	gAnimationManager = animation.NewAnimationManager()
	gAnimationManager.RegisterAnimation(
		&animation.Animation{
			Frames: []*ebiten.Image{
				charImageAsset.GetTileFromOffset(55),
				charImageAsset.GetTileFromOffset(56),
				charImageAsset.GetTileFromOffset(57),
				charImageAsset.GetTileFromOffset(56),
			},
			Delay: 8,
		},
		animation.IdPlayerIdleAnimation,
	)
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.cam.pos.X += 16
		g.player.pos.X -= 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.cam.pos.X -= 16
		g.player.pos.X += 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.cam.pos.Y += 16
		g.player.pos.Y -= 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.cam.pos.Y -= 16
		g.player.pos.Y += 1
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
	o.GeoM.Translate(g.cam.pos.X+float64(g.player.pos.X*constants.TileSize), g.cam.pos.Y+float64(g.player.pos.Y*constants.TileSize))
	screen.DrawImage(g.player.anim.GetCurFrame(), o)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	g := &Game{
		player: Player{
			pos:  vec.Vec2{X: 4, Y: 4},
			anim: gAnimationManager.GetAnimation(animation.IdPlayerIdleAnimation),
		},
		cam: Camera{
			pos: vec.Vec2{X: 0, Y: 0},
		},
		area: Area{
			name: "Test area",
			layout: []int{
				1, 1, 1, 1, 1, 1, 1, 1,
				1, 4, 4, 4, 4, 4, 4, 1,
				1, 4, 4, 4, 4, 4, 4, 1,
				1, 4, 4, 4, 4, 4, 4, 1,
				1, 1, 1, 18, 1, 1, 1, 1,
			},
			tilesW: 8,
			tilesH: 5,
		},
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
