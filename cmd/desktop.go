package main

import (
	"bytes"
	"fmt"
	"game_development/config"
	"game_development/drawing"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/spf13/viper"
)

type (
	Game struct {
		title string
		text  string
	}
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var (
	arcadeFontText *text.GoTextFaceSource
)

func init() {
	arcadeText, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.ArcadeN_ttf))
	if err != nil {
		log.Fatal(err)
	}
	arcadeFontText = arcadeText
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const (
		normalFontSize = 24
		bigFontSize    = 48
	)
	var (
		x   = 20
		msg = fmt.Sprintf("%s - TPS: %0.2f", g.text, ebiten.ActualTPS())
	)

	drawing := drawing.NewDrawing(&text.DrawOptions{})
	drawing.Upper(screen, x, msg, arcadeFontText, normalFontSize)
	drawing.Middle(screen, x, color.White, g.title, arcadeFontText, normalFontSize)
	drawing.Below(screen, x, color.White, "Main Lobby", arcadeFontText, normalFontSize)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	config.ExecConfig()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(viper.GetString("game.title"))
	if err := ebiten.RunGame(&Game{
		title: viper.GetString("game.title"),
		text:  "Welcome to Game",
	}); err != nil {
		log.Fatal(err)
	}
}
