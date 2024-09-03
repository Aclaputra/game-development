package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/Aclaputra/game-development/config"
	"github.com/Aclaputra/game-development/constant"
	"github.com/Aclaputra/game-development/drawing"
	"github.com/Aclaputra/game-development/game/npc"
	"github.com/Aclaputra/game-development/model"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/spf13/viper"
)

type (
	Game struct {
		Title string
		Text  string
	}
)

func (g *Game) Update() error {
	model.CountMovement++
	model.TimeCounter++

	skeleton := npc.NewSkeleton()
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

	drawSprite := drawing.NewDrawSprite(&ebiten.DrawImageOptions{})
	drawSprite.Position(screen, model.SkeletonSprite, float64(model.CountMovement), 500)

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
