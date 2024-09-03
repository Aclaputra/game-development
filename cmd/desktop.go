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
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/spf13/viper"
)

type (
	Game struct {
		title         string
		text          string
		countMovement int
		timeCounter   int
	}
)

const (
	SCREEN_WIDTH         = 1280
	SCREEN_HEIGHT        = 720
	RESET_FROM_START     = 0
	SKELETON_SPRITE_PATH = "assets\\lpcentry\\png\\walkcycle\\BODY_skeleton.png"
)

var (
	arcadeFontText     *text.GoTextFaceSource
	skeletonSprite     *ebiten.Image
	skeletonFrameIndex int
	skeletonFramePixel int
	skeletonStepFrames = []int{
		0,   // ok
		60,  // ok
		125, // ok
		185, // ok
		250, // ok
		310, // ok
		375, // ok
		440, // ok
		505, // ok
	}
	skeletonDirectionFrames = map[string]int{
		"east":  200,
		"south": 130,
	}
)

func (g *Game) Update() error {
	g.countMovement++
	g.timeCounter++
	fmt.Println(g.countMovement)

	skeletonFramePixel = skeletonStepFrames[skeletonFrameIndex]
	skeletonImg, err := helper.LoadAndCropImage(SKELETON_SPRITE_PATH, skeletonFramePixel, skeletonDirectionFrames["east"], 50, 75)
	if err != nil {
		return fmt.Errorf("cannot get %v", SKELETON_SPRITE_PATH)
	}
	skeletonSprite = skeletonImg

	if g.timeCounter >= 10 {
		skeletonFrameIndex++
		g.timeCounter = RESET_FROM_START
	}

	if skeletonFrameIndex >= len(skeletonStepFrames) {
		skeletonFrameIndex = RESET_FROM_START
	}

	if g.countMovement >= 1000 {
		g.countMovement = RESET_FROM_START
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
		msg = fmt.Sprintf("%s - TPS: %0.2f", g.text, ebiten.ActualTPS())
	)

	drawText := drawing.NewDrawText(&text.DrawOptions{})
	drawText.UpperHeader(screen, x, msg, arcadeFontText, normalFontSize)
	drawText.MiddleHeader(screen, 0, color.White, g.title, arcadeFontText, normalFontSize)
	drawText.BelowHeader(screen, 0, color.White, "Main Lobby", arcadeFontText, normalFontSize)

	drawSprite := drawing.NewDrawSprite(&ebiten.DrawImageOptions{})
	drawSprite.Position(screen, skeletonSprite, float64(g.countMovement), 500)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Pixel at: %v", skeletonFramePixel))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func init() {
	arcadeText, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.ArcadeN_ttf))
	if err != nil {
		log.Fatal(err)
	}
	arcadeFontText = arcadeText

}

func main() {
	config.ExecConfig()
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle(viper.GetString("game.title"))

	gameTitle := viper.GetString("game.title")
	if err := ebiten.RunGame(&Game{
		title: gameTitle,
		text:  fmt.Sprintf("Welcome to %s", gameTitle),
	}); err != nil {
		log.Fatal(err)
	}
}
