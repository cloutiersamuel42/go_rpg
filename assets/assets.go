package assets

import (
	"image"

	"github.com/cloutiersamuel42/game/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	_ = iota
	IdAssetBasictiles
	IdAssetCharacters
)

type Asset interface {
	LoadAsset() error
	PathStr() string
}

type ImageAsset struct {
	Path        string
	Image       *ebiten.Image
	TilesWidth  int
	TilesHeight int
}

func (ia *ImageAsset) LoadAsset() error {
	img, _, err := ebitenutil.NewImageFromFile(ia.Path)
	if err != nil {
		return err
	}
	ia.Image = img
	ia.TilesWidth = ia.Image.Bounds().Dx() / constants.TileSize
	ia.TilesHeight = ia.Image.Bounds().Dy() / constants.TileSize
	return nil
}

func (ia *ImageAsset) GetTileFromOffset(offset int) *ebiten.Image {
	if offset == 0 {
		return nil
	}

	realOffset := offset - 1

	tileX := realOffset % ia.TilesWidth
	tileY := realOffset / ia.TilesWidth
	posX := tileX * constants.TileSize
	posY := tileY * constants.TileSize

	return ia.Image.SubImage(image.Rect(posX, posY, posX+constants.TileSize, posY+constants.TileSize)).(*ebiten.Image)

}

func (ia *ImageAsset) PathStr() string { return ia.Path }
