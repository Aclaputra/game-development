package main

import (
	"bytes"
	"fmt"
	"game_development/drawing"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type (
	Game struct {
		title string
		text  string
	}
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
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
	drawing.Info(screen, x, msg, mplusFaceSource, normalFontSize)
	drawing.Text(screen, x, color.White, g.title, mplusFaceSource, normalFontSize)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Font (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{
		title: "Game Development",
		text:  "Welcome to Game Development",
	}); err != nil {
		log.Fatal(err)
	}
}
