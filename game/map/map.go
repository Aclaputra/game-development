package game_map

import (
	"fmt"

	"github.com/Aclaputra/game-development/helper"
	"github.com/Aclaputra/game-development/model"
)

type (
	GameMap interface {
		RenderTile(cropX int, cropY int, gameTileSpriteLocactionIdx int) error
	}
	gameMap struct{}
)

func NewGameMap() GameMap {
	return &gameMap{}
}

func (gm *gameMap) RenderTile(cropX int, cropY int, gameTileSpriteLocactionIdx int) error {
	reqLoadAndCropImage := &model.LoadAndCropImageRequest{
		Path:   "assets\\map\\PathAndObjects.png",
		X:      cropX,
		Y:      cropY,
		Width:  50,
		Height: 50,
	}
	gameMapImg, err := helper.LoadAndCropImage(reqLoadAndCropImage)
	if err != nil {
		return fmt.Errorf("cannot get %v", "assets\\map\\PathAndObjects.png")
	}
	model.GameTileSprite = append(model.GameTileSprite, gameMapImg)

	return nil
}
