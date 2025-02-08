package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/Aclaputra/game-development/entities"
	"github.com/Aclaputra/game-development/game"
	"github.com/Aclaputra/game-development/tilemap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	// scale         = 2.0
	SCREEN_WIDTH  = 320
	SCREEN_HEIGHT = 240
)

type (
	Game struct {
		Player        *entities.Player
		enemies       []*entities.Enemy
		potions       []*entities.Potion
		tilemapJSON   tilemap.TileMapJSON
		tilemapImage  *ebiten.Image
		camera        *game.Camera
		statusMessage string
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
			g.statusMessage = "Picked Up Potion"
		} else {
			g.statusMessage = ""
		}
	}

	g.camera.FollowTarget(g.Player.X+8, g.Player.Y+8, float64(SCREEN_WIDTH), float64(SCREEN_HEIGHT))
	g.camera.Constrain(
		float64(g.tilemapJSON.Layers[0].Width)*16.0,
		float64(g.tilemapJSON.Layers[0].Height)*16.0,
		float64(SCREEN_HEIGHT),
		float64(SCREEN_WIDTH),
	)

	return nil
}

// stacking dari atas ke bawah
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}

	// loop over the layers
	for _, layer := range g.tilemapJSON.Layers {
		// loop over the tiles in the layer data
		for index, id := range layer.Data {
			// get the tile position of the tile
			x := index % layer.Width
			y := index / layer.Width

			// convert the tile posititon to pixel position
			x *= 16
			y *= 16

			srcX := (id - 1) % 22
			srcY := (id - 1) / 22

			srcX *= 16
			srcY *= 16

			// opts.GeoM.Scale(scale, scale)
			opts.GeoM.Translate(float64(x), float64(y))
			opts.GeoM.Translate(g.camera.X, g.camera.Y)

			screen.DrawImage(
				g.tilemapImage.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image),
				&opts,
			)

			opts.GeoM.Reset()
		}
	}

	// opts.GeoM.Scale(scale, scale)
	opts.GeoM.Translate(g.Player.X, g.Player.Y)
	opts.GeoM.Translate(g.camera.X, g.camera.Y)

	// draw player
	screen.DrawImage(
		g.Player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()

	for _, sprite := range g.enemies {
		// opts.GeoM.Scale(scale, scale)
		opts.GeoM.Translate(sprite.X, sprite.Y)
		opts.GeoM.Translate(g.camera.X, g.camera.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

	for _, potion := range g.potions {
		// opts.GeoM.Scale(scale, scale)
		opts.GeoM.Translate(potion.X, potion.Y)
		opts.GeoM.Translate(g.camera.X, g.camera.Y)

		screen.DrawImage(
			potion.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

	ebitenutil.DebugPrint(screen, g.statusMessage)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
	// return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("RPG 2D")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	var (
		err                                            error
		playerImg, cavemenImg, potionImg, tilemapImage *ebiten.Image
		tilemapJSON                                    *tilemap.TileMapJSON
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
	if tilemapImage, _, err = ebitenutil.NewImageFromFile("assets/ninjapack/Backgrounds/Tilesets/TilesetFloor.png"); err != nil {
		log.Fatal(err)
	}
	if tilemapJSON, err = tilemap.NewTileMapJSON("assets/map/spawn.json"); err != nil {
		log.Fatal(err)
	}

	game := &Game{
		Player: &entities.Player{
			Sprite: &entities.Sprite{
				Img: playerImg,
				X:   300,
				Y:   230,
			},
			Health: 1000,
		},
		enemies: []*entities.Enemy{
			{
				Sprite: &entities.Sprite{
					Img: cavemenImg,
					X:   200,
					Y:   100,
				},
				FollowsPlayer: true,
			},
			{
				Sprite: &entities.Sprite{
					Img: cavemenImg,
					X:   250,
					Y:   100,
				},
				FollowsPlayer: false,
			},
		},
		potions: []*entities.Potion{
			{
				Sprite: &entities.Sprite{
					Img: potionImg,
					X:   320,
					Y:   120,
				},
				AmtHeal: 100,
			},
		},
		tilemapJSON:  *tilemapJSON,
		tilemapImage: tilemapImage,
		camera:       game.NewCamera(0.0, 0.0),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
