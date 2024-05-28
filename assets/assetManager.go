package assets

import (
	"fmt"
	"log"
)

type AssetManager struct {
	assets map[int]Asset
}

func NewAssetManager() *AssetManager {
	return &AssetManager{
		assets: make(map[int]Asset),
	}
}

func (am *AssetManager) RegisterAsset(asset Asset, id int) {
	am.assets[id] = asset
}

func (am *AssetManager) LoadAssets() {
	for _, asset := range am.assets {
		if err := asset.LoadAsset(); err != nil {
			log.Fatalf("Fatal error: Could not load asset: %s\n", asset.PathStr())
		} else {
			fmt.Printf("Loaded asset: %s\n", asset.PathStr())
		}
	}
	fmt.Printf("Loaded all assets!\n")
}

func (am *AssetManager) GetAsset(id int) Asset {
	asset, ok := am.assets[id]
	if !ok {
		return nil
	}
	return asset
}
