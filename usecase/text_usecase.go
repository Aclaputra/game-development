package usecase

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type (
	TextUsecase interface {
		DrawSampleText(screen *ebiten.Image, x int, msg string, mplusFaceSource *text.GoTextFaceSource, normalFontSize float64)
	}
	textUsecase struct {
		op *text.DrawOptions
	}
)

func NewTextUsecase(op *text.DrawOptions) TextUsecase {
	return &textUsecase{
		op: op,
	}
}

func (tu *textUsecase) DrawSampleText(screen *ebiten.Image, x int, msg string, mplusFaceSource *text.GoTextFaceSource, normalFontSize float64) {
	tu.op.GeoM.Translate(float64(x), 10)
	tu.op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   normalFontSize,
	}, tu.op)
}
