package main

import (
	"bytes"
	"fmt"
	"game_development/config"
	"game_development/drawing"
	"game_development/helper"
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
		count int
	}
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var (
	arcadeFontText *text.GoTextFaceSource
	ballSprite     *ebiten.Image
)

func (g *Game) Update() error {
	g.count++
	fmt.Println(g.count)
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

	drawText := drawing.NewDrawText(&text.DrawOptions{})
	drawText.UpperHeader(screen, x, msg, arcadeFontText, normalFontSize)
	drawText.MiddleHeader(screen, 0, color.White, g.title, arcadeFontText, normalFontSize)
	drawText.BelowHeader(screen, 0, color.White, "Main Lobby", arcadeFontText, normalFontSize)

	drawSprite := drawing.NewDrawSprite(&ebiten.DrawImageOptions{})
	drawSprite.Position(screen, ballSprite, float64(g.count), 500)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func init() {
	arcadeText, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.ArcadeN_ttf))
	if err != nil {
		log.Fatal(err)
	}
	arcadeFontText = arcadeText

	ballPath := "ball.png"
	ballImg, err := helper.LoadImage(ballPath)
	if err != nil {
		panic(fmt.Sprintf("cannot get %v", ballPath))
	}
	ballSprite = ballImg
}

func main() {
	config.ExecConfig()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(viper.GetString("game.title"))

	gameTitle := viper.GetString("game.title")
	if err := ebiten.RunGame(&Game{
		title: gameTitle,
		text:  fmt.Sprintf("Welcome to %s", gameTitle),
	}); err != nil {
		log.Fatal(err)
	}
}
