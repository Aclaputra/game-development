package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	scale         = 2.0
	SCREEN_WIDTH  = 320
	SCREEN_HEIGHT = 240
)

type (
	Game struct {
		Player  *Player
		enemies []*Enemy
		potions []*Potion
	}
	Sprite struct {
		Img  *ebiten.Image
		X, Y float64
	}
	Potion struct {
		*Sprite
		AmtHeal uint
	}
	// embedded struct
	Enemy struct {
		*Sprite
		FollowsPlayer bool
	}
	Player struct {
		*Sprite
		Health uint
	}
)

func (g *Game) Update() error {
	// player controller input
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.Y += 2
	}

	for _, enemy := range g.enemies {
		if enemy.FollowsPlayer {
			if enemy.X < g.Player.X {
				enemy.X += 1
			} else if enemy.X > g.Player.X {
				enemy.X -= 1
			}

			if enemy.Y < g.Player.Y {
				enemy.Y += 1
			} else if enemy.Y > g.Player.Y {
				enemy.Y -= 1
			}
		}
	}

	for _, potion := range g.potions {
		if g.Player.X > potion.X {
			g.Player.Health += potion.AmtHeal
			fmt.Printf("Picked up potion health increased by %v, current Health %v\n", potion.AmtHeal, g.Player.Health)
		}
	}

	return nil
}

// stacking dari atas ke bawah
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scale, scale)
	opts.GeoM.Translate(g.Player.X, g.Player.Y)

	// draw player
	screen.DrawImage(
		g.Player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()

	for _, sprite := range g.enemies {
		opts.GeoM.Scale(scale, scale)
		opts.GeoM.Translate(sprite.X, sprite.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

	for _, potion := range g.potions {
		opts.GeoM.Scale(scale, scale)
		opts.GeoM.Translate(potion.X, potion.Y)

		screen.DrawImage(
			potion.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// return outsideWidth, outsideHeight
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("RPG 2D")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	var (
		err                              error
		playerImg, cavemenImg, potionImg *ebiten.Image
	)

	if playerImg, _, err = ebitenutil.NewImageFromFile("assets/ninjapack/Actor/Characters/NinjaGray/SpriteSheet.png"); err != nil {
		log.Fatal(err)
	}
	if cavemenImg, _, err = ebitenutil.NewImageFromFile("assets/ninjapack/Actor/Characters/Caveman2/SpriteSheet.png"); err != nil {
		log.Fatal(err)
	}
	if potionImg, _, err = ebitenutil.NewImageFromFile("assets/ninjapack/Items/Potion/Heart.png"); err != nil {
		log.Fatal(err)
	}

	game := &Game{
		Player: &Player{
			&Sprite{
				Img: playerImg,
				X:   300,
				Y:   230,
			},
			1000,
		},
		enemies: []*Enemy{
			{
				&Sprite{
					Img: cavemenImg,
					X:   200,
					Y:   100,
				},
				true,
			},
			{
				&Sprite{
					Img: cavemenImg,
					X:   250,
					Y:   100,
				},
				false,
			},
		},
		potions: []*Potion{
			{
				&Sprite{
					Img: potionImg,
					X:   320,
					Y:   120,
				},
				100,
			},
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
