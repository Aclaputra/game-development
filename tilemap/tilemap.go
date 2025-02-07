package tilemap

import (
	"encoding/json"
	"os"
)

type TileMapLayerJSON struct {
	Data   []int `json:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

type TileMapJSON struct {
	Layers []TileMapLayerJSON `json:"layers"`
}

func NewTileMapJSON(filepath string) (*TileMapJSON, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tileMapJSON TileMapJSON
	if err = json.Unmarshal(contents, &tileMapJSON); err != nil {
		return nil, err
	}

	return &tileMapJSON, nil
}
