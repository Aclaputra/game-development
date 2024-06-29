package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	title string
	text  string
}

const (
	screenWidth  = 640
	screenHeight = 480
)

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

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
