package drawing

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type (
	Drawing interface {
		Info(screen *ebiten.Image, x int, msg string, mplusFaceSource *text.GoTextFaceSource, normalFontSize float64)
		Text(screen *ebiten.Image, x int, color color.Color, sampleText string, mplusFaceSource *text.GoTextFaceSource, fontSize float64)
	}
	drawing struct {
		op *text.DrawOptions
	}
)

func NewDrawing(op *text.DrawOptions) Drawing {
	return &drawing{
		op: op,
	}
}

func (tu *drawing) Info(screen *ebiten.Image, x int, msg string, mplusFaceSource *text.GoTextFaceSource, fontSize float64) {
	tu.op.GeoM.Translate(float64(x), 10)
	tu.op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   fontSize,
	}, tu.op)
}

func (tu *drawing) Text(screen *ebiten.Image, x int, color color.Color, sampleText string, mplusFaceSource *text.GoTextFaceSource, fontSize float64) {
	tu.op = &text.DrawOptions{}
	tu.op.GeoM.Translate(float64(x), 60)
	tu.op.ColorScale.ScaleWithColor(color)
	text.Draw(screen, sampleText, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   fontSize,
	}, tu.op)
}
