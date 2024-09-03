package drawing

import "github.com/hajimehoshi/ebiten/v2"

type (
	DrawingSprite interface {
		Position(screen *ebiten.Image, sprite *ebiten.Image, tx float64, ty float64)
	}
	drawingSprite struct {
		op *ebiten.DrawImageOptions
	}
)

func NewDrawSprite(op *ebiten.DrawImageOptions) DrawingSprite {
	return &drawingSprite{
		op: op,
	}
}

func (ds *drawingSprite) Position(screen *ebiten.Image, sprite *ebiten.Image, tx float64, ty float64) {
	ds.op.GeoM.Translate(float64(tx), float64(ty))
	screen.DrawImage(sprite, ds.op)
}
