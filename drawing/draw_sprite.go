package drawing

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type (
	DrawingText interface {
		UpperHeader(screen *ebiten.Image, x int, msg string, fontType *text.GoTextFaceSource, normalFontSize float64)
		MiddleHeader(screen *ebiten.Image, x int, color color.Color, sampleText string, fontType *text.GoTextFaceSource, fontSize float64)
		BelowHeader(screen *ebiten.Image, x int, color color.Color, sampleText string, fontType *text.GoTextFaceSource, fontSize float64)
	}
	drawingText struct {
		op *text.DrawOptions
	}
)

func NewDrawText(op *text.DrawOptions) DrawingText {
	return &drawingText{
		op: op,
	}
}

func (tu *drawingText) UpperHeader(screen *ebiten.Image, x int, msg string, fontType *text.GoTextFaceSource, fontSize float64) {
	tu.op.GeoM.Translate(float64(x), 10)
	tu.op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, msg, &text.GoTextFace{
		Source: fontType,
		Size:   fontSize,
	}, tu.op)
}

func (tu *drawingText) MiddleHeader(screen *ebiten.Image, x int, color color.Color, sampleText string, fontType *text.GoTextFaceSource, fontSize float64) {
	tu.op.GeoM.Translate(float64(x), 60)
	tu.op.ColorScale.ScaleWithColor(color)
	text.Draw(screen, sampleText, &text.GoTextFace{
		Source: fontType,
		Size:   fontSize,
	}, tu.op)
}

func (tu *drawingText) BelowHeader(screen *ebiten.Image, x int, color color.Color, sampleText string, fontType *text.GoTextFaceSource, fontSize float64) {
	tu.op.GeoM.Translate(float64(x), 60)
	tu.op.ColorScale.ScaleWithColor(color)
	text.Draw(screen, sampleText, &text.GoTextFace{
		Source: fontType,
		Size:   fontSize,
	}, tu.op)
}
