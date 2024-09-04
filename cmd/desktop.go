package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/Aclaputra/game-development/config"
	"github.com/Aclaputra/game-development/constant"
	"github.com/Aclaputra/game-development/drawing"
	game_map "github.com/Aclaputra/game-development/game/map"
	"github.com/Aclaputra/game-development/game/npc"
	"github.com/Aclaputra/game-development/model"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/spf13/viper"
)

type (
	Game struct {
		Title string
		Text  string
		keys  []ebiten.Key
	}
)

func (g *Game) Update() error {
	model.TimeCounter++

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	if len(g.keys) > 0 {
		switch g.keys[0].String() {
		case "W":
			model.SkeletonDirection = "north"
		case "A":
			model.SkeletonDirection = "west"
		case "S":
			model.SkeletonDirection = "south"
		case "D":
			model.SkeletonDirection = "east"
		}
	}
	skeleton := npc.NewSkeleton(model.SkeletonDirection)
	if err := skeleton.Render(); err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const (
		normalFontSize = 24
		bigFontSize    = 48
	)
	var (
		x   = 20
		msg = fmt.Sprintf("%s - TPS: %0.2f", g.Text, ebiten.ActualTPS())
	)

	drawText := drawing.NewDrawText(&text.DrawOptions{})
	drawText.UpperHeader(screen, x, msg, model.ArcadeFontText, normalFontSize)
	drawText.MiddleHeader(screen, 0, color.White, g.Title, model.ArcadeFontText, normalFontSize)
	drawText.BelowHeader(screen, 0, color.White, "Main Lobby", model.ArcadeFontText, normalFontSize)

	gameMap := game_map.NewGameMap()
	if err := gameMap.RenderTile(0, 0, 0); err != nil {
		fmt.Printf("error rendering tile: %v", err.Error())
	}
	if err := gameMap.RenderTile(25, 0, 1); err != nil {
		fmt.Printf("error rendering tile: %v", err.Error())
	}
	if err := gameMap.RenderTile(50, 0, 2); err != nil {
		fmt.Printf("error rendering tile: %v", err.Error())
	}
	if err := gameMap.RenderTile(50, 25, 3); err != nil {
		fmt.Printf("error rendering tile: %v", err.Error())
	}
	if err := gameMap.RenderTile(25, 25, 4); err != nil {
		fmt.Printf("error rendering tile: %v", err.Error())
	}
	if err := gameMap.RenderTile(0, 25, 5); err != nil {
		fmt.Printf("error rendering tile: %v", err.Error())
	}
	if err := gameMap.RenderTile(0, 50, 6); err != nil {
		fmt.Printf("error rendering tile: %v", err.Error())
	}
	if err := gameMap.RenderTile(25, 50, 7); err != nil {
		fmt.Printf("error rendering tile: %v", err.Error())
	}
	if err := gameMap.RenderTile(50, 50, 8); err != nil {
		fmt.Printf("error rendering tile: %v", err.Error())
	}
	drawGameMap := drawing.NewDrawSprite(&ebiten.DrawImageOptions{})
	drawGameMap.Position(screen, model.GameTileSprite[0], float64(model.BasePosition)-50, float64(model.BasePosition)-50)
	drawGameMap.Position(screen, model.GameTileSprite[1], 50, 0)
	drawGameMap.Position(screen, model.GameTileSprite[2], 50, 0)
	drawGameMap.Position(screen, model.GameTileSprite[3], 0, 50)
	drawGameMap.Position(screen, model.GameTileSprite[4], -50, 0)
	drawGameMap.Position(screen, model.GameTileSprite[5], -50, 0)
	drawGameMap.Position(screen, model.GameTileSprite[6], 0, 50)
	drawGameMap.Position(screen, model.GameTileSprite[7], 50, 0)
	drawGameMap.Position(screen, model.GameTileSprite[8], 50, 0)

	drawWalkingSkeleton := drawing.NewDrawSprite(&ebiten.DrawImageOptions{})
	drawWalkingSkeleton.Position(screen, model.SkeletonSprite, float64(model.CountMovementX), float64(model.CountMovementY))

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Pixel at: %v", model.SkeletonFramePixel))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT
}

func init() {
	arcadeText, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.ArcadeN_ttf))
	if err != nil {
		log.Fatal(err)
	}
	model.ArcadeFontText = arcadeText

}

func main() {
	config.ExecConfig()
	ebiten.SetWindowSize(constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT)
	ebiten.SetWindowTitle(viper.GetString("game.title"))

	gameTitle := viper.GetString("game.title")
	if err := ebiten.RunGame(&Game{
		Title: gameTitle,
		Text:  fmt.Sprintf("Welcome to %s", gameTitle),
	}); err != nil {
		log.Fatal(err)
	}
}
